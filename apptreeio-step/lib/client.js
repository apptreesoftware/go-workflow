"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const axios = __importStar(require("axios"));
var LogLevel;
(function (LogLevel) {
    LogLevel[LogLevel["Debug"] = 0] = "Debug";
    LogLevel[LogLevel["Info"] = 1] = "Info";
    LogLevel[LogLevel["Warn"] = 2] = "Warn";
    LogLevel[LogLevel["Error"] = 3] = "Error";
})(LogLevel = exports.LogLevel || (exports.LogLevel = {}));
class ClientImpl {
    constructor(env) {
        this.env = env;
        this.httpClient = axios.default.create({
            headers: {
                Authorization: `Bearer ${env.authorization}`
            }
        });
    }
    async LogMessage(message, level = LogLevel.Debug) {
        await this.httpClient.post(this.env.logApiUrl, {
            environment: this.env,
            message: message,
            level: level
        });
    }
    async SpawnWorkflow(workflow, input) {
        let inputData = "";
        if (typeof input == "string") {
            inputData = input;
        }
        else {
            inputData = JSON.stringify(input);
        }
        await this.httpClient.post(this.env.spawnApiUrl, {
            environment: this.env,
            workflow: workflow,
            input: inputData
        });
    }
}
exports.ClientImpl = ClientImpl;
//# sourceMappingURL=client.js.map