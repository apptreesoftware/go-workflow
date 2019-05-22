import 'dart:convert';

class LibraryPackage {
  LibraryPackage(
    this.packageName,
    this.description,
    this.version,
  );

  String packageName;
  String description;
  String version;

  factory LibraryPackage.fromJson(Map<String, dynamic> json) {
    return new LibraryPackage(
      json['packageName'] as String,
      json['description'] as String,
      json['version'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['packageName'] = packageName;
    map['description'] = description;
    map['version'] = version;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class LibraryStep {
  LibraryStep(
    this.author,
    this.packageName,
    this.name,
    this.description,
    this.sample,
    this.inputs,
    this.outputs,
  );

  String author;
  String packageName;
  String name;
  String description;
  String sample;
  List<Input> inputs;
  List<Output> outputs;

  factory LibraryStep.fromJson(Map<String, dynamic> json) {
    return new LibraryStep(
      json['author'] as String,
      json['packageName'] as String,
      json['name'] as String,
      json['description'] as String,
      json['sample'] as String,
      json['inputs'] != null
          ? (json['inputs'] as List).map((d) => new Input.fromJson(d)).toList()
          : <Input>[],
      json['outputs'] != null
          ? (json['outputs'] as List)
              .map((d) => new Output.fromJson(d))
              .toList()
          : <Output>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['author'] = author;
    map['packageName'] = packageName;
    map['name'] = name;
    map['description'] = description;
    map['sample'] = sample;
    map['inputs'] = inputs?.map((l) => l.toJson())?.toList();
    map['outputs'] = outputs?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class Input {
  Input(
    this.name,
    this.description,
    this.sample,
    this.required,
  );

  String name;
  String description;
  String sample;
  bool required;

  factory Input.fromJson(Map<String, dynamic> json) {
    return new Input(
      json['name'] as String,
      json['description'] as String,
      json['sample'] as String,
      json['required'] as bool,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['name'] = name;
    map['description'] = description;
    map['sample'] = sample;
    map['required'] = required;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class Output {
  Output(
    this.name,
    this.description,
    this.sample,
  );

  String name;
  String description;
  String sample;

  factory Output.fromJson(Map<String, dynamic> json) {
    return new Output(
      json['name'] as String,
      json['description'] as String,
      json['sample'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['name'] = name;
    map['description'] = description;
    map['sample'] = sample;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class SearchStepsRequest {
  SearchStepsRequest(
    this.terms,
  );

  String terms;

  factory SearchStepsRequest.fromJson(Map<String, dynamic> json) {
    return new SearchStepsRequest(
      json['terms'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['terms'] = terms;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class SearchStepsResponse {
  SearchStepsResponse(
    this.steps,
  );

  List<LibraryStep> steps;

  factory SearchStepsResponse.fromJson(Map<String, dynamic> json) {
    return new SearchStepsResponse(
      json['steps'] != null
          ? (json['steps'] as List)
              .map((d) => new LibraryStep.fromJson(d))
              .toList()
          : <LibraryStep>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['steps'] = steps?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetAllStepsRequest {
  GetAllStepsRequest();

  factory GetAllStepsRequest.fromJson(Map<String, dynamic> json) {
    return new GetAllStepsRequest();
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

class GetAllStepsResponse {
  GetAllStepsResponse(
    this.steps,
  );

  List<LibraryStep> steps;

  factory GetAllStepsResponse.fromJson(Map<String, dynamic> json) {
    return new GetAllStepsResponse(
      json['steps'] != null
          ? (json['steps'] as List)
              .map((d) => new LibraryStep.fromJson(d))
              .toList()
          : <LibraryStep>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['steps'] = steps?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetSingleStepRequest {
  GetSingleStepRequest(
    this.stepId,
  );

  String stepId;

  factory GetSingleStepRequest.fromJson(Map<String, dynamic> json) {
    return new GetSingleStepRequest(
      json['stepId'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['stepId'] = stepId;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetSingleStepResponse {
  GetSingleStepResponse(
    this.step,
  );

  LibraryStep step;

  factory GetSingleStepResponse.fromJson(Map<String, dynamic> json) {
    return new GetSingleStepResponse(
      new LibraryStep.fromJson(json['step']),
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

class GetStepsForPackageRequest {
  GetStepsForPackageRequest(
    this.packageName,
    this.version,
  );

  String packageName;
  String version;

  factory GetStepsForPackageRequest.fromJson(Map<String, dynamic> json) {
    return new GetStepsForPackageRequest(
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

class GetStepsForPackageResponse {
  GetStepsForPackageResponse(
    this.steps,
  );

  List<LibraryStep> steps;

  factory GetStepsForPackageResponse.fromJson(Map<String, dynamic> json) {
    return new GetStepsForPackageResponse(
      json['steps'] != null
          ? (json['steps'] as List)
              .map((d) => new LibraryStep.fromJson(d))
              .toList()
          : <LibraryStep>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['steps'] = steps?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetAllPackagesInfoRequest {
  GetAllPackagesInfoRequest();

  factory GetAllPackagesInfoRequest.fromJson(Map<String, dynamic> json) {
    return new GetAllPackagesInfoRequest();
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

class GetAllPackagesInfoResponse {
  GetAllPackagesInfoResponse(
    this.packages,
  );

  List<LibraryPackage> packages;

  factory GetAllPackagesInfoResponse.fromJson(Map<String, dynamic> json) {
    return new GetAllPackagesInfoResponse(
      json['packages'] != null
          ? (json['packages'] as List)
              .map((d) => new LibraryPackage.fromJson(d))
              .toList()
          : <LibraryPackage>[],
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['packages'] = packages?.map((l) => l.toJson())?.toList();
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetPackageInfoRequest {
  GetPackageInfoRequest(
    this.packageName,
  );

  String packageName;

  factory GetPackageInfoRequest.fromJson(Map<String, dynamic> json) {
    return new GetPackageInfoRequest(
      json['packageName'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['packageName'] = packageName;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

class GetPackageInfoResponse {
  GetPackageInfoResponse(
    this.package,
  );

  LibraryPackage package;

  factory GetPackageInfoResponse.fromJson(Map<String, dynamic> json) {
    return new GetPackageInfoResponse(
      new LibraryPackage.fromJson(json['package']),
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

class StepPublishedPackageRequest {
  StepPublishedPackageRequest(
    this.publishId,
  );

  String publishId;

  factory StepPublishedPackageRequest.fromJson(Map<String, dynamic> json) {
    return new StepPublishedPackageRequest(
      json['publishId'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['publishId'] = publishId;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}
