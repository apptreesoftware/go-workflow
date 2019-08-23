// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Defines the API for a step using the SDK to talk to an engine.
//
'use strict';
var grpc = require('grpc');
var step_pb = require('./step_pb.js');
var common_pb = require('./common_pb.js');

function serialize_core_CachePullRequest(arg) {
  if (!(arg instanceof common_pb.CachePullRequest)) {
    throw new Error('Expected argument of type core.CachePullRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_CachePullRequest(buffer_arg) {
  return common_pb.CachePullRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_CachePullResponse(arg) {
  if (!(arg instanceof common_pb.CachePullResponse)) {
    throw new Error('Expected argument of type core.CachePullResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_CachePullResponse(buffer_arg) {
  return common_pb.CachePullResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_CachePushRequest(arg) {
  if (!(arg instanceof common_pb.CachePushRequest)) {
    throw new Error('Expected argument of type core.CachePushRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_CachePushRequest(buffer_arg) {
  return common_pb.CachePushRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_CachePushResponse(arg) {
  if (!(arg instanceof common_pb.CachePushResponse)) {
    throw new Error('Expected argument of type core.CachePushResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_CachePushResponse(buffer_arg) {
  return common_pb.CachePushResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_CacheSearchRequest(arg) {
  if (!(arg instanceof common_pb.CacheSearchRequest)) {
    throw new Error('Expected argument of type core.CacheSearchRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_CacheSearchRequest(buffer_arg) {
  return common_pb.CacheSearchRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_CacheSearchResponse(arg) {
  if (!(arg instanceof common_pb.CacheSearchResponse)) {
    throw new Error('Expected argument of type core.CacheSearchResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_CacheSearchResponse(buffer_arg) {
  return common_pb.CacheSearchResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_EmptyMessage(arg) {
  if (!(arg instanceof common_pb.EmptyMessage)) {
    throw new Error('Expected argument of type core.EmptyMessage');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_EmptyMessage(buffer_arg) {
  return common_pb.EmptyMessage.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_GetWorkflowUrlRequest(arg) {
  if (!(arg instanceof common_pb.GetWorkflowUrlRequest)) {
    throw new Error('Expected argument of type core.GetWorkflowUrlRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_GetWorkflowUrlRequest(buffer_arg) {
  return common_pb.GetWorkflowUrlRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_GetWorkflowUrlResponse(arg) {
  if (!(arg instanceof common_pb.GetWorkflowUrlResponse)) {
    throw new Error('Expected argument of type core.GetWorkflowUrlResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_GetWorkflowUrlResponse(buffer_arg) {
  return common_pb.GetWorkflowUrlResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_Package(arg) {
  if (!(arg instanceof common_pb.Package)) {
    throw new Error('Expected argument of type core.Package');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_Package(buffer_arg) {
  return common_pb.Package.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_RunStepRequest(arg) {
  if (!(arg instanceof common_pb.RunStepRequest)) {
    throw new Error('Expected argument of type core.RunStepRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_RunStepRequest(buffer_arg) {
  return common_pb.RunStepRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_StepOutput(arg) {
  if (!(arg instanceof step_pb.StepOutput)) {
    throw new Error('Expected argument of type core.StepOutput');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_StepOutput(buffer_arg) {
  return step_pb.StepOutput.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_StepQueueWorkflowRequest(arg) {
  if (!(arg instanceof step_pb.StepQueueWorkflowRequest)) {
    throw new Error('Expected argument of type core.StepQueueWorkflowRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_StepQueueWorkflowRequest(buffer_arg) {
  return step_pb.StepQueueWorkflowRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_core_StepQueueWorkflowResponse(arg) {
  if (!(arg instanceof step_pb.StepQueueWorkflowResponse)) {
    throw new Error('Expected argument of type core.StepQueueWorkflowResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_core_StepQueueWorkflowResponse(buffer_arg) {
  return step_pb.StepQueueWorkflowResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var EngineStepAPIService = exports.EngineStepAPIService = {
  ping: {
    path: '/core.EngineStepAPI/Ping',
    requestStream: false,
    responseStream: false,
    requestType: common_pb.EmptyMessage,
    responseType: common_pb.EmptyMessage,
    requestSerialize: serialize_core_EmptyMessage,
    requestDeserialize: deserialize_core_EmptyMessage,
    responseSerialize: serialize_core_EmptyMessage,
    responseDeserialize: deserialize_core_EmptyMessage,
  },
  cachePush: {
    path: '/core.EngineStepAPI/CachePush',
    requestStream: false,
    responseStream: false,
    requestType: common_pb.CachePushRequest,
    responseType: common_pb.CachePushResponse,
    requestSerialize: serialize_core_CachePushRequest,
    requestDeserialize: deserialize_core_CachePushRequest,
    responseSerialize: serialize_core_CachePushResponse,
    responseDeserialize: deserialize_core_CachePushResponse,
  },
  cachePull: {
    path: '/core.EngineStepAPI/CachePull',
    requestStream: false,
    responseStream: false,
    requestType: common_pb.CachePullRequest,
    responseType: common_pb.CachePullResponse,
    requestSerialize: serialize_core_CachePullRequest,
    requestDeserialize: deserialize_core_CachePullRequest,
    responseSerialize: serialize_core_CachePullResponse,
    responseDeserialize: deserialize_core_CachePullResponse,
  },
  cacheSearch: {
    path: '/core.EngineStepAPI/CacheSearch',
    requestStream: false,
    responseStream: false,
    requestType: common_pb.CacheSearchRequest,
    responseType: common_pb.CacheSearchResponse,
    requestSerialize: serialize_core_CacheSearchRequest,
    requestDeserialize: deserialize_core_CacheSearchRequest,
    responseSerialize: serialize_core_CacheSearchResponse,
    responseDeserialize: deserialize_core_CacheSearchResponse,
  },
  queueWorkflow: {
    path: '/core.EngineStepAPI/QueueWorkflow',
    requestStream: false,
    responseStream: false,
    requestType: step_pb.StepQueueWorkflowRequest,
    responseType: step_pb.StepQueueWorkflowResponse,
    requestSerialize: serialize_core_StepQueueWorkflowRequest,
    requestDeserialize: deserialize_core_StepQueueWorkflowRequest,
    responseSerialize: serialize_core_StepQueueWorkflowResponse,
    responseDeserialize: deserialize_core_StepQueueWorkflowResponse,
  },
  getWorkflowUrl: {
    path: '/core.EngineStepAPI/GetWorkflowUrl',
    requestStream: false,
    responseStream: false,
    requestType: common_pb.GetWorkflowUrlRequest,
    responseType: common_pb.GetWorkflowUrlResponse,
    requestSerialize: serialize_core_GetWorkflowUrlRequest,
    requestDeserialize: deserialize_core_GetWorkflowUrlRequest,
    responseSerialize: serialize_core_GetWorkflowUrlResponse,
    responseDeserialize: deserialize_core_GetWorkflowUrlResponse,
  },
};

exports.EngineStepAPIClient = grpc.makeGenericClientConstructor(EngineStepAPIService);
var StepHostService = exports.StepHostService = {
  runStep: {
    path: '/core.StepHost/RunStep',
    requestStream: false,
    responseStream: false,
    requestType: common_pb.RunStepRequest,
    responseType: step_pb.StepOutput,
    requestSerialize: serialize_core_RunStepRequest,
    requestDeserialize: deserialize_core_RunStepRequest,
    responseSerialize: serialize_core_StepOutput,
    responseDeserialize: deserialize_core_StepOutput,
  },
  getPackageInfo: {
    path: '/core.StepHost/GetPackageInfo',
    requestStream: false,
    responseStream: false,
    requestType: common_pb.EmptyMessage,
    responseType: common_pb.Package,
    requestSerialize: serialize_core_EmptyMessage,
    requestDeserialize: deserialize_core_EmptyMessage,
    responseSerialize: serialize_core_Package,
    responseDeserialize: deserialize_core_Package,
  },
};

exports.StepHostClient = grpc.makeGenericClientConstructor(StepHostService);
