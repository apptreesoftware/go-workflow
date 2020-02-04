"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const client_1 = require("./client");
var client_2 = require("./client");
exports.LogLevel = client_2.LogLevel;
function createClient(data) {
    let env;
    if (isRequest(data)) {
        env = getEnvironment(data);
    }
    else {
        env = data;
    }
    return new client_1.ClientImpl(env);
}
exports.createClient = createClient;
function getEnvironment(req) {
    const env = req.body._environment;
    if (!env) {
        throw Error("Request does not contain an environment");
    }
    return env;
}
exports.getEnvironment = getEnvironment;
function validateInputs(req, ...inputs) {
    const body = req.body;
    for (let input of inputs) {
        if (!body[input]) {
            throw Error(`required input '${input}' was not provided`);
        }
    }
}
exports.validateInputs = validateInputs;
function isRequest(data) {
    if (!data) {
        return false;
    }
    if (data.protocol) {
        return true;
    }
    return false;
}
//# sourceMappingURL=index.js.map