"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
const vscode = require("vscode");
const axios = require("axios");
// this method is called when your extension is activated
// your extension is activated the very first time the command is executed
function activate(context) {
    // Use the console to output diagnostic information (console.log) and errors (console.error)
    // This line of code will only be executed once when your extension is activated
    console.log('Congratulations, your extension "apptree-steps" is now active!');
    // The command has been defined in the package.json file
    // Now provide the implementation of the command with registerCommand
    // The commandId parameter must match the command field in package.json
    let findStep = vscode.commands.registerCommand('extension.findStep', () => __awaiter(this, void 0, void 0, function* () {
        // The code you place here will be executed every time your command is executed
        // Display a message box to the user
        var packages = yield getAvailablePackages();
        var pkg = yield vscode.window.showQuickPick(packages);
        if (pkg === undefined) {
            return;
        }
        var steps = yield getStepsForPackage(pkg);
        if (steps === undefined) {
            vscode.window.showErrorMessage("No steps found");
        }
        var items = steps.map((s) => {
            return { label: s.name, description: s.author, data: s };
        });
        var step = yield vscode.window.showQuickPick(items);
        if (step === undefined) {
            return;
        }
        const yaml = yield getYamlForStep(pkg, step.label);
        insertStep(yaml);
    }));
    context.subscriptions.push(findStep);
}
exports.activate = activate;
// this method is called when your extension is deactivated
function deactivate() { }
exports.deactivate = deactivate;
function getAvailablePackages() {
    return __awaiter(this, void 0, void 0, function* () {
        let resp = yield axios.default.get('https://platform.apptreeio.com/api/packages');
        return resp.data;
    });
}
function getStepsForPackage(pkg) {
    return __awaiter(this, void 0, void 0, function* () {
        let resp = yield axios.default.get(`https://platform.apptreeio.com/api/packages/${pkg}/steps`);
        return resp.data;
    });
}
function getYamlForStep(pkg, step) {
    return __awaiter(this, void 0, void 0, function* () {
        let resp = yield axios.default.get(`https://platform.apptreeio.com/api/packages/${pkg}/steps/${step}/yaml`);
        return resp.data;
    });
}
function insertStep(text) {
    let editor = vscode.window.activeTextEditor;
    if (editor === undefined) {
        return;
    }
    let activeEditor = editor;
    let snippet = new vscode.SnippetString(`- ${text}`);
    editor.insertSnippet(snippet);
}
//# sourceMappingURL=extension.js.map