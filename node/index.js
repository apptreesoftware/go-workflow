#!/usr/bin/env node
'use strict';
const service = require('./gen/step_grpc_pb');
const yaml = require('yaml');
const grpc = require('@grpc/grpc-js');
const fs = require("fs");
const steps = {};
const yargs = require('yargs');
const getStdin = require('get-stdin');

module.exports.runId = process.env['RUN_ID'];
module.exports.EXIT_WORKFLOW_SUCCESS_CODE = 200;
module.exports.EXIT_WORKFLOW_FAIL_CODE = 201;

module.exports.addStep = function (name, version, func) {
    steps[`${name}@${version}`] = func;
};

module.exports.validateInputs = function(...inputNames) {
    let inputArray = [];
    if (Array.isArray(inputNames)) {
        inputArray = inputNames;
    } else {
        inputArray = [inputNames];
    }
    for (let name of inputArray) {
        let val = module.exports.stepInput[name];
        if (val == null) {
            throw `Input '${name}' is required but not defined`;
        }
    }
};

module.exports.run = async function () {
    const args = process.argv;
    const serveMode = args.includes("--serve");
    const portIndex = args.indexOf("--port");
    let port = "4000";
    if (portIndex !== -1 && args.length > portIndex + 1) {
        port = args[portIndex + 1];
    }
    if (serveMode) {
        _runServeMode(port);
    } else {
        try {
            await _runIPCMode();
            process.exit(0);
        } catch (e) {
            console.error(e);
            process.exit(1);
        }
    }
};

function _runServeMode(port) {
    const server = new grpc.Server();
    server.addService(service.StepHostService, { runStep: runStep, getPackageInfo: getPackageInfo });
    server.bindAsync(`0.0.0.0:${port}`, grpc.ServerCredentials.createInsecure(), function (err) {
        if (err === null) {
            server.start();
        }
    });
    console.log(`Starting step package on port ${port}`);

}

async function _runIPCMode() {
    const env = process.env;
    let name = env['STEP_NAME'];
    let version = env['STEP_VERSION'];
    const outputPath = env['WORKFLOW_OUTPUT'];
    module.exports.debug = env['DEBUG'] === 'true';
    if (version === undefined || version == null) {
        version = '1.0';
    }
    const args = yargs.argv;
    if (args['step'] != null) {
        name = args['step'];
    }
    if (args['debug'] === true) {
        module.exports.debug = true
    }
    const stepId = `${name}@${version}`;
    const step = steps[stepId];
    if (step == null) {
        console.error(`${stepId} not found. Make sure this step is registered using 'addStep'`);
    }

    let fileName = args['file'];
    let input = null;
    if (fileName != null) {
        input = fs.readFileSync(fileName, 'utf8');
    }
    let sampleFileName = args['samples'];
    if (sampleFileName != null) {
        try {
            const sampleInput = fs.readFileSync(sampleFileName, 'utf8');
            const sampleJson = JSON.parse(sampleInput);
            let stepInputObject = null;
            if (sampleJson[name] != null) {
                stepInputObject = sampleJson[name];
            } else if (sampleJson[stepId] != null) {
                stepInputObject = sampleJson[stepId];
            }
            if (stepInputObject == null) {
                console.error(`The sample file does not contain an entry for '${name}' or '${stepId}'`);
                return;
            }
            input = JSON.stringify(stepInputObject);
        } catch (e) {
            console.error(`A samples file was provided but data for step ${stepId} was not found. Make sure the file ${sampleFileName} exists and contains an entry for ${stepId}`)
            return;
        }
    }

    if (input == null) {
        let dataStr = await getStdin();
        await _runIPCWithInput(step, dataStr, outputPath);
    } else {
        await _runIPCWithInput(step, input, outputPath);
    }
}

async function _runIPCWithInput(step, input, outputPath) {
    try {
        const parsedData = JSON.parse(input);
        module.exports.stepInput = parsedData;
        let respOrPromise = step(parsedData);
        let resp;
        if (respOrPromise instanceof Promise) {
            resp = await respOrPromise;
        } else {
            resp = respOrPromise;
        }
        const outData = JSON.stringify(resp);
        if (outputPath === undefined || outputPath === null) {
            process.stdout.write(outData);
        } else {
            fs.writeFileSync(outputPath, outData);
        }
    } catch (e) {
        console.error(e);
        throw `Could not write step output: ${e}`
    }
}

async function runStep(call, callback) {
    const reply = new proto.core.StepOutput();
    const input = call.request.getInput();
    const env = call.request.getEnvironment();
    const stepId = `${env.getStepname()}@${env.getStepversion()}`;
    const step = steps[stepId];
    const dec = new TextDecoder();
    const enc = new TextEncoder();
    try {
        if (step == null) {
            callback({message: "step not found", status: grpc.status.NOT_FOUND});
            return;
        }
        const inputStr = dec.decode(input);
        const inputJson = JSON.parse(inputStr);
        module.exports.stepInput = inputJson;
        let resp = step(inputJson);
        if (resp instanceof Promise) {
            resp = await resp;
        }
        const respBytes = enc.encode(JSON.stringify(resp));

        reply.setOutput(respBytes);
        callback(null, reply);
    } catch (e) {
        console.error(e);
        callback({code: 400, message: e.toString(), status: grpc.status.INTERNAL}, null);
    }
}

function getPackageInfo(call, callback) {
    const file = fs.readFileSync('./package.yaml', 'utf8');
    const parsed = yaml.parse(file);
    const pkg = new proto.core.Package();

    pkg.setName(parsed['name']);
    pkg.setVersion(`${parsed['version']}`);
    pkg.setLang(parsed['lang']);

    const map = pkg.getStepsMap();
    Object.keys(parsed.steps).forEach((k) => {
        let stepDef = parsed.steps[k];
        let step = new proto.core.PackageStep();
        step.description = stepDef['description'];
        step.version = stepDef['version'];
        let stepInputMap = step.getInputsMap();

        // Get inputs
        if (stepDef.inputs != null) {
            Object.keys(stepDef.inputs).forEach((inputKey) => {
                let inputDef = stepDef.inputs[inputKey];
                let input = new proto.core.InputInfo();
                input.setRequired(inputDef.required);
                input.setDescription(inputDef.description);
                input.setSample(inputDef.sample);
                stepInputMap.set(inputKey, input);
            });
        }
        // Get outputs

        let stepOutputMap = step.getOutputsMap();
        if (stepDef.outputs != null) {
            Object.keys(stepDef.outputs).forEach((outputKey) => {
                let outputDef = stepDef.outputs[outputKey];
                let output = new proto.core.OutputInfo();
                output.setDescription(outputDef['description']);
                stepOutputMap.set(outputKey, output);
            });
        }
        map.set(k, step);
    });

    callback(null, pkg);
}