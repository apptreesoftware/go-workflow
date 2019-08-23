#!/usr/bin/env node
'use strict';
const service = require('./gen/step_grpc_pb');
var core = require('./gen/common_pb');
const yaml = require('yaml');
const grpc = require('grpc');
const fs = require("fs");
const steps = {};

module.exports.addStep = function (name, version, func) {
    steps[`${name}-${version}`] = func;
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
        _runIPCMode();
    }
};

function _runServeMode(port) {
    const server = new grpc.Server();
    server.addService(service.StepHostService, { runStep: runStep, getPackageInfo: getPackageInfo });
    server.bind(`0.0.0.0:${port}`, grpc.ServerCredentials.createInsecure());
    console.log(`Starting step package on port ${port}`);
    server.start();
}

function _runIPCMode() {
    const env = process.env;
    const name = env['STEP_NAME'];
    const version = env['STEP_VERSION'];
    const outputPath = env['WORKFLOW_OUTPUT'];
    module.exports.debug = env['DEBUG'] == 'true';
    const stepId = `${name}-${version}`;
    const step = steps[stepId];
    if (step == null) {
        console.error(`${stepId} not found. Make sure this step is registered using 'addStep'`);
    }
    let dataStr = "";
    process.stdin.setEncoding("utf8");
    process.stdin.on('data', (data) => {
        console.log(`Input: ${data}`);
        dataStr += data;
    });
    process.stdin.on('end', () => {
        _runIPCWithInput(step, dataStr, outputPath);
    });
}

function _runIPCWithInput(step, input, outputPath) {
    try {
        const parsedData = JSON.parse(input);
        const resp = step(parsedData);

        const outData = JSON.stringify(resp);
        if (outputPath === undefined || outputPath === null) {
            process.stdout.write(outData);
        } else {
            fs.writeFile(outputPath, outData, (err) => {
                if (err) {
                    console.log(err);
                }
            });
        }
    }
    catch (e) {
        console.error(e);
    }
}

function runStep(call, callback) {
    const reply = new proto.core.StepOutput();
    const input = call.request.getInput();
    const env = call.request.getEnvironment();
    const stepId = `${env.getStepname()}-${env.getStepversion()}`;
    const step = steps[stepId];
    const dec = new TextDecoder();
    const enc = new TextEncoder();
    try {
        if (step == null) {
            callback({ message: "step not found", status: grpc.status.NOT_FOUND });
            return;
        }
        const inputStr = dec.decode(input);
        const inputJson = JSON.parse(inputStr);
        const resp = step(inputJson);
        const respBytes = enc.encode(JSON.stringify(resp));

        reply.setOutput(respBytes);
        callback(null, reply);
    } catch (e) {
        console.error(e);
        callback({ code: 400, message: e.toString(), status: grpc.status.INTERNAL }, null);
    }
}

function getPackageInfo(call, callback) {
    const file = fs.readFileSync('./package.yaml', 'utf8')
    const parsed = yaml.parse(file);
    console.log(parsed);
    var pkg = new core.Package();

    pkg.setName(parsed['name']);
    pkg.setVersion(`${parsed['version']}`);
    pkg.setLang(parsed['lang']);

    const map = pkg.getStepsMap();
    Object.keys(parsed.steps).forEach((k) => {
        let stepDef = parsed.steps[k];
        let step = new core.PackageStep();
        step.description = stepDef['description'];
        step.version = stepDef['version'];
        let stepInputMap = step.getInputsMap();

        // Get inputs
        if (stepDef.inputs != null) {
            Object.keys(stepDef.inputs).forEach((inputKey) => {
                let inputDef = stepDef.inputs[inputKey];
                let input = new core.InputInfo();
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
                let output = new core.OutputInfo();
                output.setDescription(outputDef['description']);
                stepOutputMap.set(outputKey, output);
            });
        }
        map.set(k, step);
    });

    callback(null, pkg);
}
