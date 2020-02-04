import * as express from "express";
import { Environment, ClientImpl, Client } from "./client";

export { Environment, Client, LogLevel } from "./client";

export function createClient(data: express.Request | Environment): Client {
  let env: Environment;
  if (isRequest(data)) {
    env = getEnvironment(data as express.Request);
  } else {
    env = data as Environment;
  }
  return new ClientImpl(env);
}

export function getEnvironment(req: express.Request): Environment {
  const env = req.body._environment as Environment;
  if (!env) {
    throw Error("Request does not contain an environment");
  }
  return env;
}

export function validateInputs(req: express.Request, ...inputs: string[]) {
  const body = req.body;
  for (let input of inputs) {
    if (!body[input]) {
      throw Error(`required input '${input}' was not provided`);
    }
  }
}

function isRequest(data: any): boolean {
  if (!data) {
    return false;
  }
  if (data.protocol) {
    return true;
  }
  return false;
}
