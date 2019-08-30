// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from 'vscode';
import * as axios from 'axios';


// this method is called when your extension is activated
// your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {

	// Use the console to output diagnostic information (console.log) and errors (console.error)
	// This line of code will only be executed once when your extension is activated
	console.log('Congratulations, your extension "apptree-steps" is now active!');

	// The command has been defined in the package.json file
	// Now provide the implementation of the command with registerCommand
	// The commandId parameter must match the command field in package.json

	let findStep = vscode.commands.registerCommand('extension.findStep', async () => {
		// The code you place here will be executed every time your command is executed

		// Display a message box to the user
		var packages = await getAvailablePackages();
		var pkg = await vscode.window.showQuickPick(packages);
		if (pkg === undefined) {
			return;
		}
		var steps = await getStepsForPackage(pkg);
		if (steps === undefined) {
			vscode.window.showErrorMessage("No steps found");
		}

		var items: vscode.QuickPickItem[] = steps.map((s) => {
			return { label: s.name, description: s.author, data: s };
		});

		var step = await vscode.window.showQuickPick(items);
		if (step === undefined) {
			return;
		}

		const yaml = await getYamlForStep(pkg, step.label);
		insertStep(yaml);
	});
	context.subscriptions.push(findStep);
}

// this method is called when your extension is deactivated
export function deactivate() { }

async function getAvailablePackages(): Promise<string[]> {
	let resp = await axios.default.get('https://platform.apptreeio.com/api/packages');
	return resp.data;
}

async function getStepsForPackage(pkg: string): Promise<Step[]> {
	let resp = await axios.default.get(`https://platform.apptreeio.com/api/packages/${pkg}/steps`);
	return resp.data as Step[];
}

async function getYamlForStep(pkg: string, step: string): Promise<string> {
	let resp = await axios.default.get(`https://platform.apptreeio.com/api/packages/${pkg}/steps/${step}/yaml`);
	return resp.data;
}

function insertStep(text: string) {
	let editor = vscode.window.activeTextEditor;
	if (editor === undefined) {
		return;
	}
	let activeEditor = editor;
	let snippet = new vscode.SnippetString(`- ${text}`);

	editor.insertSnippet(snippet);
}

export interface Input {
	name: string;
	description: string;
	sample: string;
	required: boolean;
}

export interface Output {
	name: string;
	description: string;
}

export interface Step {
	author: string;
	packageName: string;
	name: string;
	inputs: Input[];
	outputs: Output[];
}