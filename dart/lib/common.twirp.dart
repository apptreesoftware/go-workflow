import 'dart:convert';

class Package {
  Package(
    this.name,
    this.lang,
    this.version,
    this.executables,
    this.steps,
    this.description,
  );

  String name;
  String lang;
  String version;
  Exec executables;
  Map<String, PackageStep> steps;
  String description;

  factory Package.fromJson(Map<String, dynamic> json) {
    var stepsMap = new Map<String, PackageStep>();
    (json['steps'] as Map<String, dynamic>)?.forEach((key, val) {
      stepsMap[key] = new PackageStep.fromJson(val as Map<String, dynamic>);
    });

    return new Package(
      json['name'] as String,
      json['lang'] as String,
      json['version'] as String,
      new Exec.fromJson(json['executables']),
      stepsMap,
      json['description'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['name'] = name;
    map['lang'] = lang;
    map['version'] = version;
    map['executables'] = executables.toJson();
    map['steps'] = json.decode(json.encode(steps));
    map['description'] = description;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class Exec {
  Exec(
    this.darwin,
    this.linux,
    this.windows,
  );

  Binary darwin;
  Binary linux;
  Binary windows;

  factory Exec.fromJson(Map<String, dynamic> json) {
    return new Exec(
      new Binary.fromJson(json['darwin']),
      new Binary.fromJson(json['linux']),
      new Binary.fromJson(json['windows']),
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['darwin'] = darwin.toJson();
    map['linux'] = linux.toJson();
    map['windows'] = windows.toJson();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class Binary {
  Binary(
    this.amd64,
    this.i386,
  );

  String amd64;
  String i386;

  factory Binary.fromJson(Map<String, dynamic> json) {
    return new Binary(
      json['amd64'] as String,
      json['i386'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['amd64'] = amd64;
    map['i386'] = i386;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class PackageStep {
  PackageStep(
    this.description,
    this.inputs,
    this.outputs,
    this.sample,
  );

  String description;
  Map<String, InputInfo> inputs;
  Map<String, OutputInfo> outputs;
  String sample;

  factory PackageStep.fromJson(Map<String, dynamic> json) {
    var inputsMap = new Map<String, InputInfo>();
    (json['inputs'] as Map<String, dynamic>)?.forEach((key, val) {
      inputsMap[key] = new InputInfo.fromJson(val as Map<String, dynamic>);
    });

    var outputsMap = new Map<String, OutputInfo>();
    (json['outputs'] as Map<String, dynamic>)?.forEach((key, val) {
      outputsMap[key] = new OutputInfo.fromJson(val as Map<String, dynamic>);
    });

    return new PackageStep(
      json['description'] as String,
      inputsMap,
      outputsMap,
      json['sample'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['description'] = description;
    map['inputs'] = json.decode(json.encode(inputs));
    map['outputs'] = json.decode(json.encode(outputs));
    map['sample'] = sample;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class InputInfo {
  InputInfo(
    this.required,
    this.description,
    this.sample,
  );

  bool required;
  String description;
  String sample;

  factory InputInfo.fromJson(Map<String, dynamic> json) {
    return new InputInfo(
      json['required'] as bool,
      json['description'] as String,
      json['sample'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['required'] = required;
    map['description'] = description;
    map['sample'] = sample;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class OutputInfo {
  OutputInfo(
    this.description,
  );

  String description;

  factory OutputInfo.fromJson(Map<String, dynamic> json) {
    return new OutputInfo(
      json['description'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['description'] = description;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class Environment {
  Environment(
    this.project,
    this.workflow,
    this.runId,
    this.stepName,
    this.stepVersion,
    this.inputFile,
    this.triggerType,
    this.stepInstanceId,
    this.package,
    this.cacheHost,
    this.allowParallel,
    this.debug,
  );

  String project;
  String workflow;
  String runId;
  String stepName;
  String stepVersion;
  String inputFile;
  String triggerType;
  String stepInstanceId;
  String package;
  String cacheHost;
  bool allowParallel;
  bool debug;

  factory Environment.fromJson(Map<String, dynamic> json) {
    return new Environment(
      json['project'] as String,
      json['workflow'] as String,
      json['runId'] as String,
      json['stepName'] as String,
      json['stepVersion'] as String,
      json['inputFile'] as String,
      json['triggerType'] as String,
      json['stepInstanceId'] as String,
      json['package'] as String,
      json['cacheHost'] as String,
      json['allowParallel'] as bool,
      json['debug'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    map['workflow'] = workflow;
    map['runId'] = runId;
    map['stepName'] = stepName;
    map['stepVersion'] = stepVersion;
    map['inputFile'] = inputFile;
    map['triggerType'] = triggerType;
    map['stepInstanceId'] = stepInstanceId;
    map['package'] = package;
    map['cacheHost'] = cacheHost;
    map['allowParallel'] = allowParallel;
    map['debug'] = debug;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class EmptyMessage {
  EmptyMessage();

  factory EmptyMessage.fromJson(Map<String, dynamic> json) {
    return new EmptyMessage();
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class StepPackageRequest {
  StepPackageRequest(
    this.id,
    this.version,
    this.os,
    this.arch,
  );

  String id;
  String version;
  String os;
  String arch;

  factory StepPackageRequest.fromJson(Map<String, dynamic> json) {
    return new StepPackageRequest(
      json['id'] as String,
      json['version'] as String,
      json['os'] as String,
      json['arch'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['id'] = id;
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

class StepPackageIdRequest {
  StepPackageIdRequest(
    this.id,
  );

  String id;

  factory StepPackageIdRequest.fromJson(Map<String, dynamic> json) {
    return new StepPackageIdRequest(
      json['id'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['id'] = id;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class StepPackageResponse {
  StepPackageResponse(
    this.packageUrl,
    this.publishId,
  );

  String packageUrl;
  String publishId;

  factory StepPackageResponse.fromJson(Map<String, dynamic> json) {
    return new StepPackageResponse(
      json['packageUrl'] as String,
      json['publishId'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['packageUrl'] = packageUrl;
    map['publishId'] = publishId;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetStepPackageRequest {
  GetStepPackageRequest(
    this.packageName,
    this.version,
  );

  String packageName;
  String version;

  factory GetStepPackageRequest.fromJson(Map<String, dynamic> json) {
    return new GetStepPackageRequest(
      json['packageName'] as String,
      json['version'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['packageName'] = packageName;
    map['version'] = version;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetStepPackageResponse {
  GetStepPackageResponse(
    this.package,
  );

  Package package;

  factory GetStepPackageResponse.fromJson(Map<String, dynamic> json) {
    return new GetStepPackageResponse(
      new Package.fromJson(json['package']),
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['package'] = package.toJson();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class CachePushRequest {
  CachePushRequest(
    this.id,
    this.metadata,
    this.record,
    this.cacheName,
    this.environment,
  );

  String id;
  String metadata;
  String record;
  String cacheName;
  Environment environment;

  factory CachePushRequest.fromJson(Map<String, dynamic> json) {
    return new CachePushRequest(
      json['id'] as String,
      json['metadata'] as String,
      json['record'] as String,
      json['cacheName'] as String,
      new Environment.fromJson(json['environment']),
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['id'] = id;
    map['metadata'] = metadata;
    map['record'] = record;
    map['cacheName'] = cacheName;
    map['environment'] = environment.toJson();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class CachePushResponse {
  CachePushResponse();

  factory CachePushResponse.fromJson(Map<String, dynamic> json) {
    return new CachePushResponse();
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class CachePullRequest {
  CachePullRequest(
    this.id,
    this.cacheName,
    this.environment,
  );

  String id;
  String cacheName;
  Environment environment;

  factory CachePullRequest.fromJson(Map<String, dynamic> json) {
    return new CachePullRequest(
      json['id'] as String,
      json['cacheName'] as String,
      new Environment.fromJson(json['environment']),
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['id'] = id;
    map['cacheName'] = cacheName;
    map['environment'] = environment.toJson();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class CachePullResponse {
  CachePullResponse(
    this.record,
    this.metadata,
    this.notFound,
  );

  String record;
  String metadata;
  bool notFound;

  factory CachePullResponse.fromJson(Map<String, dynamic> json) {
    return new CachePullResponse(
      json['record'] as String,
      json['metadata'] as String,
      json['notFound'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['record'] = record;
    map['metadata'] = metadata;
    map['notFound'] = notFound;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class CacheSearchRequest {
  CacheSearchRequest(
    this.cacheName,
    this.searchFilter,
    this.environment,
    this.limit,
  );

  String cacheName;
  String searchFilter;
  Environment environment;
  int limit;

  factory CacheSearchRequest.fromJson(Map<String, dynamic> json) {
    return new CacheSearchRequest(
      json['cacheName'] as String,
      json['searchFilter'] as String,
      new Environment.fromJson(json['environment']),
      json['limit'] as int,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['cacheName'] = cacheName;
    map['searchFilter'] = searchFilter;
    map['environment'] = environment.toJson();
    map['limit'] = limit;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class CacheSearchResponse {
  CacheSearchResponse(
    this.records,
  );

  List<RawRecord> records;

  factory CacheSearchResponse.fromJson(Map<String, dynamic> json) {
    return new CacheSearchResponse(
      json['records'] != null
          ? (json['records'] as List)
              .map((d) => new RawRecord.fromJson(d))
              .toList()
          : <RawRecord>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['records'] = records?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RawRecord {
  RawRecord(
    this.record,
    this.metadata,
  );

  String record;
  String metadata;

  factory RawRecord.fromJson(Map<String, dynamic> json) {
    return new RawRecord(
      json['record'] as String,
      json['metadata'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['record'] = record;
    map['metadata'] = metadata;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class AllStepsRequest {
  AllStepsRequest(
    this.searchTerm,
  );

  String searchTerm;

  factory AllStepsRequest.fromJson(Map<String, dynamic> json) {
    return new AllStepsRequest(
      json['searchTerm'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['searchTerm'] = searchTerm;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class AllStepsResponse {
  AllStepsResponse(
    this.success,
    this.message,
    this.recordCount,
    this.items,
  );

  bool success;
  String message;
  int recordCount;
  List<RegisteredStep> items;

  factory AllStepsResponse.fromJson(Map<String, dynamic> json) {
    return new AllStepsResponse(
      json['success'] as bool,
      json['message'] as String,
      json['recordCount'] as int,
      json['items'] != null
          ? (json['items'] as List)
              .map((d) => new RegisteredStep.fromJson(d))
              .toList()
          : <RegisteredStep>[],
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

class RegisteredStep {
  RegisteredStep(
    this.step,
    this.location,
    this.locationType,
    this.publishId,
    this.author,
    this.packageName,
    this.StepName,
  );

  PackageStep step;
  String location;
  String locationType;
  String publishId;
  String author;
  String packageName;
  String StepName;

  factory RegisteredStep.fromJson(Map<String, dynamic> json) {
    return new RegisteredStep(
      new PackageStep.fromJson(json['step']),
      json['location'] as String,
      json['locationType'] as String,
      json['publishId'] as String,
      json['author'] as String,
      json['packageName'] as String,
      json['StepName'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['step'] = step.toJson();
    map['location'] = location;
    map['locationType'] = locationType;
    map['publishId'] = publishId;
    map['author'] = author;
    map['packageName'] = packageName;
    map['StepName'] = StepName;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class SingleStepRequest {
  SingleStepRequest(
    this.packageName,
    this.stepName,
    this.version,
  );

  String packageName;
  String stepName;
  String version;

  factory SingleStepRequest.fromJson(Map<String, dynamic> json) {
    return new SingleStepRequest(
      json['packageName'] as String,
      json['stepName'] as String,
      json['version'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['packageName'] = packageName;
    map['stepName'] = stepName;
    map['version'] = version;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class SingleStepResponse {
  SingleStepResponse(
    this.step,
  );

  RegisteredStep step;

  factory SingleStepResponse.fromJson(Map<String, dynamic> json) {
    return new SingleStepResponse(
      new RegisteredStep.fromJson(json['step']),
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['step'] = step.toJson();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class RunStepRequest {
  RunStepRequest(
    this.environment,
    this.input,
    this.stepConfig,
  );

  Environment environment;
  String input;
  String stepConfig;

  factory RunStepRequest.fromJson(Map<String, dynamic> json) {
    return new RunStepRequest(
      new Environment.fromJson(json['environment']),
      json['input'] as String,
      json['stepConfig'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['environment'] = environment.toJson();
    map['input'] = input;
    map['stepConfig'] = stepConfig;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ProjectRequest {
  ProjectRequest(
    this.project,
  );

  String project;

  factory ProjectRequest.fromJson(Map<String, dynamic> json) {
    return new ProjectRequest(
      json['project'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['project'] = project;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class UserRequest {
  UserRequest(
    this.username,
  );

  String username;

  factory UserRequest.fromJson(Map<String, dynamic> json) {
    return new UserRequest(
      json['username'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['username'] = username;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class BasicResponse {
  BasicResponse(
    this.success,
    this.message,
  );

  bool success;
  String message;

  factory BasicResponse.fromJson(Map<String, dynamic> json) {
    return new BasicResponse(
      json['success'] as bool,
      json['message'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['success'] = success;
    map['message'] = message;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class Empty {
  Empty();

  factory Empty.fromJson(Map<String, dynamic> json) {
    return new Empty();
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class AllPackagesNamesResponse {
  AllPackagesNamesResponse(
    this.packages,
  );

  List<String> packages;

  factory AllPackagesNamesResponse.fromJson(Map<String, dynamic> json) {
    return new AllPackagesNamesResponse(
      json['packages'] != null
          ? (json['packages'] as List).cast<String>()
          : <String>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['packages'] = packages?.map((l) => l)?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class ListProjectResponse {
  ListProjectResponse(
    this.success,
    this.message,
    this.projects,
  );

  bool success;
  String message;
  List<String> projects;

  factory ListProjectResponse.fromJson(Map<String, dynamic> json) {
    return new ListProjectResponse(
      json['success'] as bool,
      json['message'] as String,
      json['projects'] != null
          ? (json['projects'] as List).cast<String>()
          : <String>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['success'] = success;
    map['message'] = message;
    map['projects'] = projects?.map((l) => l)?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetWorkflowUrlRequest {
  GetWorkflowUrlRequest(
    this.environment,
    this.workflow,
    this.params,
  );

  Environment environment;
  String workflow;
  Map<String, String> params;

  factory GetWorkflowUrlRequest.fromJson(Map<String, dynamic> json) {
    var paramsMap = new Map<String, String>();
    (json['params'] as Map<String, dynamic>)?.forEach((key, val) {
      if (val is String) {
      } else if (val is num) {}
    });

    return new GetWorkflowUrlRequest(
      new Environment.fromJson(json['environment']),
      json['workflow'] as String,
      paramsMap,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['environment'] = environment.toJson();
    map['workflow'] = workflow;
    map['params'] = json.decode(json.encode(params));
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetWorkflowUrlResponse {
  GetWorkflowUrlResponse(
    this.workflowUrl,
  );

  String workflowUrl;

  factory GetWorkflowUrlResponse.fromJson(Map<String, dynamic> json) {
    return new GetWorkflowUrlResponse(
      json['workflowUrl'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['workflowUrl'] = workflowUrl;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}
