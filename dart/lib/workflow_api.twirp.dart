import 'dart:async';
import 'package:http/http.dart';
import 'package:requester/requester.dart';
import 'package:twirp_dart_core/twirp_dart_core.dart';
import 'dart:convert';
import 'common.twirp.dart';

class ProjectWorkflowRequest {
  ProjectWorkflowRequest(
    this.project,
    this.workflow,
  );

  String project;
  String workflow;

  factory ProjectWorkflowRequest.fromJson(Map<String, dynamic> json) {
    return new ProjectWorkflowRequest(
      json['project'] as String,
      json['workflow'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['workflow'] = workflow;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class StoreValueRequest {
  StoreValueRequest(
    this.key,
    this.value,
    this.project,
    this.isSecret,
  );

  String key;
  String value;
  String project;
  bool isSecret;

  factory StoreValueRequest.fromJson(Map<String, dynamic> json) {
    return new StoreValueRequest(
      json['key'] as String,
      json['value'] as String,
      json['project'] as String,
      json['isSecret'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['key'] = key;
    map['value'] = value;
    map['project'] = project;
    map['isSecret'] = isSecret;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class DeleteValueRequest {
  DeleteValueRequest(
    this.key,
    this.project,
  );

  String key;
  String project;

  factory DeleteValueRequest.fromJson(Map<String, dynamic> json) {
    return new DeleteValueRequest(
      json['key'] as String,
      json['project'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['key'] = key;
    map['project'] = project;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class SingleValueResponse {
  SingleValueResponse(
    this.value,
  );

  String value;

  factory SingleValueResponse.fromJson(Map<String, dynamic> json) {
    return new SingleValueResponse(
      json['value'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['value'] = value;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ValueResponse {
  ValueResponse(
    this.values,
  );

  Map<String, String> values;

  factory ValueResponse.fromJson(Map<String, dynamic> json) {
    var valuesMap = new Map<String, String>();
    (json['values'] as Map<String, dynamic>)?.forEach((key, val) {
      if (val is String) {
      } else if (val is num) {}
    });

    return new ValueResponse(
      valuesMap,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['values'] = json.decode(json.encode(values));
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ProjectValue {
  ProjectValue(
    this.key,
    this.value,
    this.project,
    this.isEncrypted,
  );

  String key;
  String value;
  String project;
  bool isEncrypted;

  factory ProjectValue.fromJson(Map<String, dynamic> json) {
    return new ProjectValue(
      json['key'] as String,
      json['value'] as String,
      json['project'] as String,
      json['isEncrypted'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['key'] = key;
    map['value'] = value;
    map['project'] = project;
    map['isEncrypted'] = isEncrypted;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class JobSearch {
  JobSearch(
    this.project,
    this.workflow,
    this.jobId,
    this.spawnedFromJobId,
    this.status,
    this.fromDate,
    this.toDate,
    this.count,
    this.offset,
  );

  String project;
  String workflow;
  String jobId;
  String spawnedFromJobId;
  String status;
  int fromDate;
  int toDate;
  int count;
  int offset;

  factory JobSearch.fromJson(Map<String, dynamic> json) {
    return new JobSearch(
      json['project'] as String,
      json['workflow'] as String,
      json['jobId'] as String,
      json['spawnedFromJobId'] as String,
      json['status'] as String,
      json['fromDate'] as int,
      json['toDate'] as int,
      json['count'] as int,
      json['offset'] as int,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['workflow'] = workflow;
    map['jobId'] = jobId;
    map['spawnedFromJobId'] = spawnedFromJobId;
    map['status'] = status;
    map['fromDate'] = fromDate;
    map['toDate'] = toDate;
    map['count'] = count;
    map['offset'] = offset;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class JobSearchResponse {
  JobSearchResponse(
    this.items,
    this.moreAvailable,
  );

  List<JobSearchResult> items;
  bool moreAvailable;

  factory JobSearchResponse.fromJson(Map<String, dynamic> json) {
    return new JobSearchResponse(
      json['items'] != null
          ? (json['items'] as List)
              .map((d) => new JobSearchResult.fromJson(d))
              .toList()
          : <JobSearchResult>[],
      json['moreAvailable'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['items'] = items?.map((l) => l.toJson())?.toList();
    map['moreAvailable'] = moreAvailable;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class JobSearchResult {
  JobSearchResult(
    this.id,
    this.success,
    this.start,
    this.end,
    this.duration,
    this.failureInfo,
    this.project,
    this.workflow,
    this.triggerType,
    this.input,
    this.spawnCount,
    this.spawnTrail,
    this.statusMessage,
    this.steps,
  );

  String id;
  bool success;
  int start;
  int end;
  double duration;
  JobFailureInfo failureInfo;
  String project;
  String workflow;
  String triggerType;
  String input;
  int spawnCount;
  List<SpawnTrailEntry> spawnTrail;
  String statusMessage;
  List<JobStepSummary> steps;

  factory JobSearchResult.fromJson(Map<String, dynamic> json) {
    return new JobSearchResult(
      json['id'] as String,
      json['success'] as bool,
      json['start'] as int,
      json['end'] as int,
      json['duration'] as double,
      new JobFailureInfo.fromJson(json['failureInfo']),
      json['project'] as String,
      json['workflow'] as String,
      json['triggerType'] as String,
      json['input'] as String,
      json['spawnCount'] as int,
      json['spawnTrail'] != null
          ? (json['spawnTrail'] as List)
              .map((d) => new SpawnTrailEntry.fromJson(d))
              .toList()
          : <SpawnTrailEntry>[],
      json['statusMessage'] as String,
      json['steps'] != null
          ? (json['steps'] as List)
              .map((d) => new JobStepSummary.fromJson(d))
              .toList()
          : <JobStepSummary>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['id'] = id;
    map['success'] = success;
    map['start'] = start;
    map['end'] = end;
    map['duration'] = duration;
    map['failureInfo'] = failureInfo.toJson();
    map['project'] = project;
    map['workflow'] = workflow;
    map['triggerType'] = triggerType;
    map['input'] = input;
    map['spawnCount'] = spawnCount;
    map['spawnTrail'] = spawnTrail?.map((l) => l.toJson())?.toList();
    map['statusMessage'] = statusMessage;
    map['steps'] = steps?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class JobStepSummary {
  JobStepSummary(
    this.stepId,
    this.duration,
    this.status,
    this.message,
  );

  String stepId;
  double duration;
  String status;
  String message;

  factory JobStepSummary.fromJson(Map<String, dynamic> json) {
    return new JobStepSummary(
      json['stepId'] as String,
      json['duration'] as double,
      json['status'] as String,
      json['message'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['stepId'] = stepId;
    map['duration'] = duration;
    map['status'] = status;
    map['message'] = message;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class JobFailureInfo {
  JobFailureInfo(
    this.stepIndex,
    this.message,
    this.stepInput,
    this.stepName,
  );

  int stepIndex;
  String message;
  String stepInput;
  String stepName;

  factory JobFailureInfo.fromJson(Map<String, dynamic> json) {
    return new JobFailureInfo(
      json['stepIndex'] as int,
      json['message'] as String,
      json['stepInput'] as String,
      json['stepName'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['stepIndex'] = stepIndex;
    map['message'] = message;
    map['stepInput'] = stepInput;
    map['stepName'] = stepName;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RemoteEngineUpdateRequest {
  RemoteEngineUpdateRequest(
    this.name,
    this.project,
    this.forceUpdate,
    this.waitForVersion,
  );

  String name;
  String project;
  bool forceUpdate;
  bool waitForVersion;

  factory RemoteEngineUpdateRequest.fromJson(Map<String, dynamic> json) {
    return new RemoteEngineUpdateRequest(
      json['name'] as String,
      json['project'] as String,
      json['forceUpdate'] as bool,
      json['waitForVersion'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['name'] = name;
    map['project'] = project;
    map['forceUpdate'] = forceUpdate;
    map['waitForVersion'] = waitForVersion;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RemoteEngineUpdateResponse {
  RemoteEngineUpdateResponse(
    this.status,
    this.version,
    this.os,
    this.arch,
  );

  String status;
  String version;
  String os;
  String arch;

  factory RemoteEngineUpdateResponse.fromJson(Map<String, dynamic> json) {
    return new RemoteEngineUpdateResponse(
      json['status'] as String,
      json['version'] as String,
      json['os'] as String,
      json['arch'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['status'] = status;
    map['version'] = version;
    map['os'] = os;
    map['arch'] = arch;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RemoteEngineHealthRequest {
  RemoteEngineHealthRequest(
    this.name,
    this.project,
  );

  String name;
  String project;

  factory RemoteEngineHealthRequest.fromJson(Map<String, dynamic> json) {
    return new RemoteEngineHealthRequest(
      json['name'] as String,
      json['project'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['name'] = name;
    map['project'] = project;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RemoteEngineHealthResponse {
  RemoteEngineHealthResponse(
    this.status,
    this.version,
    this.os,
    this.arch,
  );

  String status;
  String version;
  String os;
  String arch;

  factory RemoteEngineHealthResponse.fromJson(Map<String, dynamic> json) {
    return new RemoteEngineHealthResponse(
      json['status'] as String,
      json['version'] as String,
      json['os'] as String,
      json['arch'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['status'] = status;
    map['version'] = version;
    map['os'] = os;
    map['arch'] = arch;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ListRemoteEnginesResponse {
  ListRemoteEnginesResponse(
    this.engines,
  );

  List<RemoteEngine> engines;

  factory ListRemoteEnginesResponse.fromJson(Map<String, dynamic> json) {
    return new ListRemoteEnginesResponse(
      json['engines'] != null
          ? (json['engines'] as List)
              .map((d) => new RemoteEngine.fromJson(d))
              .toList()
          : <RemoteEngine>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['engines'] = engines?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RemoteEngine {
  RemoteEngine(
    this.name,
    this.host,
    this.apiKey,
    this.cert,
  );

  String name;
  String host;
  String apiKey;
  String cert;

  factory RemoteEngine.fromJson(Map<String, dynamic> json) {
    return new RemoteEngine(
      json['name'] as String,
      json['host'] as String,
      json['apiKey'] as String,
      json['cert'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['name'] = name;
    map['host'] = host;
    map['apiKey'] = apiKey;
    map['cert'] = cert;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RemoteEngineRequest {
  RemoteEngineRequest(
    this.project,
    this.engine,
  );

  String project;
  String engine;

  factory RemoteEngineRequest.fromJson(Map<String, dynamic> json) {
    return new RemoteEngineRequest(
      json['project'] as String,
      json['engine'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['engine'] = engine;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RegisterRemoteEngineRequest {
  RegisterRemoteEngineRequest(
    this.name,
    this.host,
    this.clientCert,
    this.apiKey,
    this.project,
  );

  String name;
  String host;
  String clientCert;
  String apiKey;
  String project;

  factory RegisterRemoteEngineRequest.fromJson(Map<String, dynamic> json) {
    return new RegisterRemoteEngineRequest(
      json['name'] as String,
      json['host'] as String,
      json['clientCert'] as String,
      json['apiKey'] as String,
      json['project'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['name'] = name;
    map['host'] = host;
    map['clientCert'] = clientCert;
    map['apiKey'] = apiKey;
    map['project'] = project;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ProjectValueRequest {
  ProjectValueRequest(
    this.project,
    this.key,
  );

  String project;
  String key;

  factory ProjectValueRequest.fromJson(Map<String, dynamic> json) {
    return new ProjectValueRequest(
      json['project'] as String,
      json['key'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['key'] = key;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ProjectValuesRequest {
  ProjectValuesRequest(
    this.project,
    this.decrypt,
  );

  String project;
  bool decrypt;

  factory ProjectValuesRequest.fromJson(Map<String, dynamic> json) {
    return new ProjectValuesRequest(
      json['project'] as String,
      json['decrypt'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['decrypt'] = decrypt;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ListWorkflowResponse {
  ListWorkflowResponse(
    this.workflows,
  );

  List<WorkflowOverview> workflows;

  factory ListWorkflowResponse.fromJson(Map<String, dynamic> json) {
    return new ListWorkflowResponse(
      json['workflows'] != null
          ? (json['workflows'] as List)
              .map((d) => new WorkflowOverview.fromJson(d))
              .toList()
          : <WorkflowOverview>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['workflows'] = workflows?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class WorkflowOverview {
  WorkflowOverview(
    this.id,
    this.project,
    this.dependencies,
    this.triggers,
  );

  String id;
  String project;
  String dependencies;
  String triggers;

  factory WorkflowOverview.fromJson(Map<String, dynamic> json) {
    return new WorkflowOverview(
      json['id'] as String,
      json['project'] as String,
      json['dependencies'] as String,
      json['triggers'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['id'] = id;
    map['project'] = project;
    map['dependencies'] = dependencies;
    map['triggers'] = triggers;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RunWorkflowRequest {
  RunWorkflowRequest(
    this.projectId,
    this.workflowId,
    this.triggerBody,
    this.count,
  );

  String projectId;
  String workflowId;
  String triggerBody;
  int count;

  factory RunWorkflowRequest.fromJson(Map<String, dynamic> json) {
    return new RunWorkflowRequest(
      json['projectId'] as String,
      json['workflowId'] as String,
      json['triggerBody'] as String,
      json['count'] as int,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['projectId'] = projectId;
    map['workflowId'] = workflowId;
    map['triggerBody'] = triggerBody;
    map['count'] = count;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RetryWorkflowRequest {
  RetryWorkflowRequest(
    this.project,
    this.workflow,
    this.fromDate,
    this.toDate,
  );

  String project;
  String workflow;
  int fromDate;
  int toDate;

  factory RetryWorkflowRequest.fromJson(Map<String, dynamic> json) {
    return new RetryWorkflowRequest(
      json['project'] as String,
      json['workflow'] as String,
      json['fromDate'] as int,
      json['toDate'] as int,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['workflow'] = workflow;
    map['fromDate'] = fromDate;
    map['toDate'] = toDate;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RetryWorkflowResponse {
  RetryWorkflowResponse(
    this.count,
    this.jobIds,
    this.errors,
  );

  int count;
  List<String> jobIds;
  List<String> errors;

  factory RetryWorkflowResponse.fromJson(Map<String, dynamic> json) {
    return new RetryWorkflowResponse(
      json['count'] as int,
      json['jobIds'] != null
          ? (json['jobIds'] as List).cast<String>()
          : <String>[],
      json['errors'] != null
          ? (json['errors'] as List).cast<String>()
          : <String>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['count'] = count;
    map['jobIds'] = jobIds?.map((l) => l)?.toList();
    map['errors'] = errors?.map((l) => l)?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RetryJobRequest {
  RetryJobRequest(
    this.project,
    this.jobId,
    this.stepIndex,
  );

  String project;
  String jobId;
  int stepIndex;

  factory RetryJobRequest.fromJson(Map<String, dynamic> json) {
    return new RetryJobRequest(
      json['project'] as String,
      json['jobId'] as String,
      json['stepIndex'] as int,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['jobId'] = jobId;
    map['stepIndex'] = stepIndex;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class PublishWorkflowRequest {
  PublishWorkflowRequest(
    this.projectId,
    this.data,
  );

  String projectId;
  String data;

  factory PublishWorkflowRequest.fromJson(Map<String, dynamic> json) {
    return new PublishWorkflowRequest(
      json['projectId'] as String,
      json['data'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['projectId'] = projectId;
    map['data'] = data;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class AddUserRequest {
  AddUserRequest(
    this.username,
    this.projectId,
  );

  String username;
  String projectId;

  factory AddUserRequest.fromJson(Map<String, dynamic> json) {
    return new AddUserRequest(
      json['username'] as String,
      json['projectId'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['username'] = username;
    map['projectId'] = projectId;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class QueueItem {
  QueueItem(
    this.project,
    this.workflow,
    this.runContext,
    this.running,
    this.createdOn,
    this.triggeredBy,
  );

  String project;
  String workflow;
  Map<String, String> runContext;
  bool running;
  String createdOn;
  String triggeredBy;

  factory QueueItem.fromJson(Map<String, dynamic> json) {
    var runContextMap = new Map<String, String>();
    (json['runContext'] as Map<String, dynamic>)?.forEach((key, val) {
      if (val is String) {
      } else if (val is num) {}
    });

    return new QueueItem(
      json['project'] as String,
      json['workflow'] as String,
      runContextMap,
      json['running'] as bool,
      json['createdOn'] as String,
      json['triggeredBy'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['workflow'] = workflow;
    map['runContext'] = json.decode(json.encode(runContext));
    map['running'] = running;
    map['createdOn'] = createdOn;
    map['triggeredBy'] = triggeredBy;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class QueueResponse {
  QueueResponse(
    this.success,
    this.message,
    this.recordCount,
    this.items,
  );

  bool success;
  String message;
  int recordCount;
  List<QueueItem> items;

  factory QueueResponse.fromJson(Map<String, dynamic> json) {
    return new QueueResponse(
      json['success'] as bool,
      json['message'] as String,
      json['recordCount'] as int,
      json['items'] != null
          ? (json['items'] as List)
              .map((d) => new QueueItem.fromJson(d))
              .toList()
          : <QueueItem>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['success'] = success;
    map['message'] = message;
    map['recordCount'] = recordCount;
    map['items'] = items?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ViewQueueRequest {
  ViewQueueRequest(
    this.project,
    this.workflow,
    this.triggerType,
  );

  String project;
  String workflow;
  String triggerType;

  factory ViewQueueRequest.fromJson(Map<String, dynamic> json) {
    return new ViewQueueRequest(
      json['project'] as String,
      json['workflow'] as String,
      json['triggerType'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['workflow'] = workflow;
    map['triggerType'] = triggerType;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ViewLogRequest {
  ViewLogRequest(
    this.project,
    this.jobId,
    this.logLevel,
  );

  String project;
  String jobId;
  String logLevel;

  factory ViewLogRequest.fromJson(Map<String, dynamic> json) {
    return new ViewLogRequest(
      json['project'] as String,
      json['jobId'] as String,
      json['logLevel'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['jobId'] = jobId;
    map['logLevel'] = logLevel;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class Log {
  Log(
    this.logLevel,
    this.message,
    this.timestamp,
  );

  String logLevel;
  String message;
  String timestamp;

  factory Log.fromJson(Map<String, dynamic> json) {
    return new Log(
      json['logLevel'] as String,
      json['message'] as String,
      json['timestamp'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['logLevel'] = logLevel;
    map['message'] = message;
    map['timestamp'] = timestamp;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ViewLogResponse {
  ViewLogResponse(
    this.success,
    this.message,
    this.logs,
  );

  bool success;
  String message;
  List<Log> logs;

  factory ViewLogResponse.fromJson(Map<String, dynamic> json) {
    return new ViewLogResponse(
      json['success'] as bool,
      json['message'] as String,
      json['logs'] != null
          ? (json['logs'] as List).map((d) => new Log.fromJson(d)).toList()
          : <Log>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['success'] = success;
    map['message'] = message;
    map['logs'] = logs?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ViewJobDetailRequest {
  ViewJobDetailRequest(
    this.project,
    this.job,
  );

  String project;
  String job;

  factory ViewJobDetailRequest.fromJson(Map<String, dynamic> json) {
    return new ViewJobDetailRequest(
      json['project'] as String,
      json['job'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['job'] = job;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ViewJobDetailResponse {
  ViewJobDetailResponse(
    this.success,
    this.message,
    this.jobDetail,
  );

  bool success;
  String message;
  JobDetail jobDetail;

  factory ViewJobDetailResponse.fromJson(Map<String, dynamic> json) {
    return new ViewJobDetailResponse(
      json['success'] as bool,
      json['message'] as String,
      new JobDetail.fromJson(json['jobDetail']),
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['success'] = success;
    map['message'] = message;
    map['jobDetail'] = jobDetail.toJson();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class JobDetail {
  JobDetail(
    this.startTime,
    this.endTime,
    this.duration,
    this.storage,
    this.stepResults,
  );

  String startTime;
  String endTime;
  double duration;
  Map<String, String> storage;
  List<StepResult> stepResults;

  factory JobDetail.fromJson(Map<String, dynamic> json) {
    var storageMap = new Map<String, String>();
    (json['storage'] as Map<String, dynamic>)?.forEach((key, val) {
      if (val is String) {
      } else if (val is num) {}
    });

    return new JobDetail(
      json['startTime'] as String,
      json['endTime'] as String,
      json['duration'] as double,
      storageMap,
      json['stepResults'] != null
          ? (json['stepResults'] as List)
              .map((d) => new StepResult.fromJson(d))
              .toList()
          : <StepResult>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['startTime'] = startTime;
    map['endTime'] = endTime;
    map['duration'] = duration;
    map['storage'] = json.decode(json.encode(storage));
    map['stepResults'] = stepResults?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class StepResult {
  StepResult(
    this.stepId,
    this.duration,
    this.status,
    this.message,
    this.inputs,
    this.outputs,
  );

  String stepId;
  int duration;
  String status;
  String message;
  Map<String, String> inputs;
  Map<String, String> outputs;

  factory StepResult.fromJson(Map<String, dynamic> json) {
    var inputsMap = new Map<String, String>();
    (json['inputs'] as Map<String, dynamic>)?.forEach((key, val) {
      if (val is String) {
      } else if (val is num) {}
    });

    var outputsMap = new Map<String, String>();
    (json['outputs'] as Map<String, dynamic>)?.forEach((key, val) {
      if (val is String) {
      } else if (val is num) {}
    });

    return new StepResult(
      json['stepId'] as String,
      json['duration'] as int,
      json['status'] as String,
      json['message'] as String,
      inputsMap,
      outputsMap,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['stepId'] = stepId;
    map['duration'] = duration;
    map['status'] = status;
    map['message'] = message;
    map['inputs'] = json.decode(json.encode(inputs));
    map['outputs'] = json.decode(json.encode(outputs));
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class TriggerResponse {
  TriggerResponse(
    this.success,
    this.message,
    this.project,
    this.jobId,
  );

  bool success;
  String message;
  String project;
  String jobId;

  factory TriggerResponse.fromJson(Map<String, dynamic> json) {
    return new TriggerResponse(
      json['success'] as bool,
      json['message'] as String,
      json['project'] as String,
      json['jobId'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['success'] = success;
    map['message'] = message;
    map['project'] = project;
    map['jobId'] = jobId;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class CancelJobRequest {
  CancelJobRequest(
    this.project,
    this.workflow,
    this.jobId,
    this.allPending,
    this.allRunning,
  );

  String project;
  String workflow;
  String jobId;
  bool allPending;
  bool allRunning;

  factory CancelJobRequest.fromJson(Map<String, dynamic> json) {
    return new CancelJobRequest(
      json['project'] as String,
      json['workflow'] as String,
      json['jobId'] as String,
      json['allPending'] as bool,
      json['allRunning'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['workflow'] = workflow;
    map['jobId'] = jobId;
    map['allPending'] = allPending;
    map['allRunning'] = allRunning;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class CancelJobResponse {
  CancelJobResponse(
    this.success,
    this.message,
    this.canceledCount,
  );

  bool success;
  String message;
  int canceledCount;

  factory CancelJobResponse.fromJson(Map<String, dynamic> json) {
    return new CancelJobResponse(
      json['success'] as bool,
      json['message'] as String,
      json['canceledCount'] as int,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['success'] = success;
    map['message'] = message;
    map['canceledCount'] = canceledCount;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class AddEventRequest {
  AddEventRequest(
    this.project,
    this.onStart,
    this.onFailure,
    this.onSuccess,
    this.dailySummary,
    this.url,
    this.extraHeaders,
  );

  String project;
  bool onStart;
  bool onFailure;
  bool onSuccess;
  bool dailySummary;
  String url;
  Map<String, String> extraHeaders;

  factory AddEventRequest.fromJson(Map<String, dynamic> json) {
    var extraHeadersMap = new Map<String, String>();
    (json['extraHeaders'] as Map<String, dynamic>)?.forEach((key, val) {
      if (val is String) {
      } else if (val is num) {}
    });

    return new AddEventRequest(
      json['project'] as String,
      json['onStart'] as bool,
      json['onFailure'] as bool,
      json['onSuccess'] as bool,
      json['dailySummary'] as bool,
      json['url'] as String,
      extraHeadersMap,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['onStart'] = onStart;
    map['onFailure'] = onFailure;
    map['onSuccess'] = onSuccess;
    map['dailySummary'] = dailySummary;
    map['url'] = url;
    map['extraHeaders'] = json.decode(json.encode(extraHeaders));
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RemoveEventRequest {
  RemoveEventRequest(
    this.project,
    this.url,
  );

  String project;
  String url;

  factory RemoveEventRequest.fromJson(Map<String, dynamic> json) {
    return new RemoveEventRequest(
      json['project'] as String,
      json['url'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['url'] = url;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

abstract class WorkflowAPI {
  Future<TriggerResponse> runWorkflow(RunWorkflowRequest runWorkflowRequest);
  Future<RetryWorkflowResponse> retryWorkflow(
      RetryWorkflowRequest retryWorkflowRequest);
  Future<TriggerResponse> retryJob(RetryJobRequest retryJobRequest);
  Future<BasicResponse> ping(EmptyMessage emptyMessage);
  Future<BasicResponse> publishWorkflow(
      PublishWorkflowRequest publishWorkflowRequest);
  Future<ListWorkflowResponse> listWorkflows(ProjectRequest projectRequest);
  Future<BasicResponse> createProject(ProjectRequest projectRequest);
  Future<BasicResponse> addUserToProject(AddUserRequest addUserRequest);
  Future<BasicResponse> addSuperUser(UserRequest userRequest);
  Future<ListProjectResponse> listProjects(Empty empty);
  Future<ListRemoteEnginesResponse> listRemoteEngines(
      ProjectRequest projectRequest);
  Future<BasicResponse> registerRemoteEngine(
      RegisterRemoteEngineRequest registerRemoteEngineRequest);
  Future<RemoteEngine> getRemoteEngine(RemoteEngineRequest remoteEngineRequest);
  Future<RemoteEngineHealthResponse> remoteEngineHealthCheck(
      RemoteEngineHealthRequest remoteEngineHealthRequest);
  Future<RemoteEngineUpdateResponse> updateRemoteEngine(
      RemoteEngineUpdateRequest remoteEngineUpdateRequest);
  Future<BasicResponse> removeRemoteEngine(
      RemoteEngineRequest remoteEngineRequest);
  Future<StepPackageResponse> getStepPackage(
      StepPackageRequest stepPackageRequest);
  Future<ViewLogResponse> viewJobLog(ViewLogRequest viewLogRequest);
  Future<ViewJobDetailResponse> viewJobDetail(
      ViewJobDetailRequest viewJobDetailRequest);
  Future<JobSearchResponse> searchJobs(JobSearch jobSearch);
  Future<BasicResponse> storeValue(StoreValueRequest storeValueRequest);
  Future<BasicResponse> deleteValue(DeleteValueRequest deleteValueRequest);
  Future<ValueResponse> getValues(ProjectValuesRequest projectValuesRequest);
  Future<SingleValueResponse> getValue(ProjectValueRequest projectValueRequest);
  Future<QueueResponse> viewQueue(ViewQueueRequest viewQueueRequest);
  Future<CancelJobResponse> cancelJobs(CancelJobRequest cancelJobRequest);
  Future<BasicResponse> disableWorkflow(
      ProjectWorkflowRequest projectWorkflowRequest);
  Future<BasicResponse> enableWorkflow(
      ProjectWorkflowRequest projectWorkflowRequest);
  Future<BasicResponse> removeWorkflow(
      ProjectWorkflowRequest projectWorkflowRequest);
  Future<BasicResponse> pauseEngines(Empty empty);
  Future<BasicResponse> unpauseEngines(Empty empty);
  Future<BasicResponse> addEvent(AddEventRequest addEventRequest);
  Future<BasicResponse> removeEvent(RemoveEventRequest removeEventRequest);
}

class DefaultWorkflowAPI implements WorkflowAPI {
  final String hostname;
  Requester _requester;
  final _pathPrefix = "/twirp/core.WorkflowAPI/";

  DefaultWorkflowAPI(this.hostname, {Requester requester}) {
    if (requester == null) {
      _requester = new Requester(new Client());
    } else {
      _requester = requester;
    }
  }

  Future<TriggerResponse> runWorkflow(
      RunWorkflowRequest runWorkflowRequest) async {
    var url = "${hostname}${_pathPrefix}RunWorkflow";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(runWorkflowRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return TriggerResponse.fromJson(value);
  }

  Future<RetryWorkflowResponse> retryWorkflow(
      RetryWorkflowRequest retryWorkflowRequest) async {
    var url = "${hostname}${_pathPrefix}RetryWorkflow";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(retryWorkflowRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return RetryWorkflowResponse.fromJson(value);
  }

  Future<TriggerResponse> retryJob(RetryJobRequest retryJobRequest) async {
    var url = "${hostname}${_pathPrefix}RetryJob";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(retryJobRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return TriggerResponse.fromJson(value);
  }

  Future<BasicResponse> ping(EmptyMessage emptyMessage) async {
    var url = "${hostname}${_pathPrefix}Ping";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(emptyMessage.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> publishWorkflow(
      PublishWorkflowRequest publishWorkflowRequest) async {
    var url = "${hostname}${_pathPrefix}PublishWorkflow";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(publishWorkflowRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<ListWorkflowResponse> listWorkflows(
      ProjectRequest projectRequest) async {
    var url = "${hostname}${_pathPrefix}ListWorkflows";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(projectRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return ListWorkflowResponse.fromJson(value);
  }

  Future<BasicResponse> createProject(ProjectRequest projectRequest) async {
    var url = "${hostname}${_pathPrefix}CreateProject";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(projectRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> addUserToProject(AddUserRequest addUserRequest) async {
    var url = "${hostname}${_pathPrefix}AddUserToProject";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(addUserRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> addSuperUser(UserRequest userRequest) async {
    var url = "${hostname}${_pathPrefix}AddSuperUser";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(userRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<ListProjectResponse> listProjects(Empty empty) async {
    var url = "${hostname}${_pathPrefix}ListProjects";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(empty.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return ListProjectResponse.fromJson(value);
  }

  Future<ListRemoteEnginesResponse> listRemoteEngines(
      ProjectRequest projectRequest) async {
    var url = "${hostname}${_pathPrefix}ListRemoteEngines";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(projectRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return ListRemoteEnginesResponse.fromJson(value);
  }

  Future<BasicResponse> registerRemoteEngine(
      RegisterRemoteEngineRequest registerRemoteEngineRequest) async {
    var url = "${hostname}${_pathPrefix}RegisterRemoteEngine";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(registerRemoteEngineRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<RemoteEngine> getRemoteEngine(
      RemoteEngineRequest remoteEngineRequest) async {
    var url = "${hostname}${_pathPrefix}GetRemoteEngine";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(remoteEngineRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return RemoteEngine.fromJson(value);
  }

  Future<RemoteEngineHealthResponse> remoteEngineHealthCheck(
      RemoteEngineHealthRequest remoteEngineHealthRequest) async {
    var url = "${hostname}${_pathPrefix}RemoteEngineHealthCheck";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(remoteEngineHealthRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return RemoteEngineHealthResponse.fromJson(value);
  }

  Future<RemoteEngineUpdateResponse> updateRemoteEngine(
      RemoteEngineUpdateRequest remoteEngineUpdateRequest) async {
    var url = "${hostname}${_pathPrefix}UpdateRemoteEngine";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(remoteEngineUpdateRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return RemoteEngineUpdateResponse.fromJson(value);
  }

  Future<BasicResponse> removeRemoteEngine(
      RemoteEngineRequest remoteEngineRequest) async {
    var url = "${hostname}${_pathPrefix}RemoveRemoteEngine";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(remoteEngineRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<StepPackageResponse> getStepPackage(
      StepPackageRequest stepPackageRequest) async {
    var url = "${hostname}${_pathPrefix}GetStepPackage";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(stepPackageRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return StepPackageResponse.fromJson(value);
  }

  Future<ViewLogResponse> viewJobLog(ViewLogRequest viewLogRequest) async {
    var url = "${hostname}${_pathPrefix}ViewJobLog";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(viewLogRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return ViewLogResponse.fromJson(value);
  }

  Future<ViewJobDetailResponse> viewJobDetail(
      ViewJobDetailRequest viewJobDetailRequest) async {
    var url = "${hostname}${_pathPrefix}ViewJobDetail";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(viewJobDetailRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return ViewJobDetailResponse.fromJson(value);
  }

  Future<JobSearchResponse> searchJobs(JobSearch jobSearch) async {
    var url = "${hostname}${_pathPrefix}SearchJobs";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(jobSearch.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return JobSearchResponse.fromJson(value);
  }

  Future<BasicResponse> storeValue(StoreValueRequest storeValueRequest) async {
    var url = "${hostname}${_pathPrefix}StoreValue";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(storeValueRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> deleteValue(
      DeleteValueRequest deleteValueRequest) async {
    var url = "${hostname}${_pathPrefix}DeleteValue";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(deleteValueRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<ValueResponse> getValues(
      ProjectValuesRequest projectValuesRequest) async {
    var url = "${hostname}${_pathPrefix}GetValues";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(projectValuesRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return ValueResponse.fromJson(value);
  }

  Future<SingleValueResponse> getValue(
      ProjectValueRequest projectValueRequest) async {
    var url = "${hostname}${_pathPrefix}GetValue";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(projectValueRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return SingleValueResponse.fromJson(value);
  }

  Future<QueueResponse> viewQueue(ViewQueueRequest viewQueueRequest) async {
    var url = "${hostname}${_pathPrefix}ViewQueue";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(viewQueueRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return QueueResponse.fromJson(value);
  }

  Future<CancelJobResponse> cancelJobs(
      CancelJobRequest cancelJobRequest) async {
    var url = "${hostname}${_pathPrefix}CancelJobs";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(cancelJobRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return CancelJobResponse.fromJson(value);
  }

  Future<BasicResponse> disableWorkflow(
      ProjectWorkflowRequest projectWorkflowRequest) async {
    var url = "${hostname}${_pathPrefix}DisableWorkflow";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(projectWorkflowRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> enableWorkflow(
      ProjectWorkflowRequest projectWorkflowRequest) async {
    var url = "${hostname}${_pathPrefix}EnableWorkflow";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(projectWorkflowRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> removeWorkflow(
      ProjectWorkflowRequest projectWorkflowRequest) async {
    var url = "${hostname}${_pathPrefix}RemoveWorkflow";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(projectWorkflowRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> pauseEngines(Empty empty) async {
    var url = "${hostname}${_pathPrefix}PauseEngines";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(empty.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> unpauseEngines(Empty empty) async {
    var url = "${hostname}${_pathPrefix}UnpauseEngines";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(empty.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> addEvent(AddEventRequest addEventRequest) async {
    var url = "${hostname}${_pathPrefix}AddEvent";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(addEventRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Future<BasicResponse> removeEvent(
      RemoveEventRequest removeEventRequest) async {
    var url = "${hostname}${_pathPrefix}RemoveEvent";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(removeEventRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return BasicResponse.fromJson(value);
  }

  Exception twirpException(Response response) {
    try {
      var value = json.decode(response.body);
      return new TwirpJsonException.fromJson(value);
    } catch (e) {
      return new TwirpException(response.body);
    }
  }
}
