import 'dart:async';
import 'package:http/http.dart';
import 'package:requester/requester.dart';
import 'package:twirp_dart_core/twirp_dart_core.dart';
import 'dart:convert';
import 'common.twirp.dart';
import 'step_library.twirp.dart';

class RegisteredPackage {
  RegisteredPackage(
    this.packageDefinition,
    this.locationType,
    this.location,
    this.publishId,
    this.author,
  );

  Package packageDefinition;
  String locationType;
  String location;
  String publishId;
  String author;

  factory RegisteredPackage.fromJson(Map<String, dynamic> json) {
    return new RegisteredPackage(
      new Package.fromJson(json['packageDefinition']),
      json['locationType'] as String,
      json['location'] as String,
      json['publishId'] as String,
      json['author'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['packageDefinition'] = packageDefinition.toJson();
    map['locationType'] = locationType;
    map['location'] = location;
    map['publishId'] = publishId;
    map['author'] = author;
    return map;
  }

  @override
  String toString() {
    return json.encode(toJson());
  }
}

abstract class StepPackageStore {
  Future<BasicResponse> ping(EmptyMessage emptyMessage);
  Future<Empty> registerPackage(RegisteredPackage registeredPackage);
  Future<RegisteredPackage> getPackage(StepPackageRequest stepPackageRequest);
  Future<RegisteredPackage> getPackageById(
      StepPackageIdRequest stepPackageIdRequest);
  Future<GetSingleStepResponse> getStep(
      GetSingleStepRequest getSingleStepRequest);
  Future<SearchStepsResponse> searchStepsForLibrary(
      SearchStepsRequest searchStepsRequest);
  Future<GetAllStepsResponse> getAllStepsForLibrary(
      GetAllStepsRequest getAllStepsRequest);
  Future<GetStepsForPackageResponse> getPackageSteps(
      GetStepsForPackageRequest getStepsForPackageRequest);
  Future<GetAllPackagesInfoResponse> getAllPackageNamesForLibrary(
      GetAllPackagesInfoRequest getAllPackagesInfoRequest);
  Future<GetPackageInfoResponse> getPackageInfoForLibrary(
      GetPackageInfoRequest getPackageInfoRequest);
}

class DefaultStepPackageStore implements StepPackageStore {
  final String hostname;
  Requester _requester;
  final _pathPrefix = "/twirp/core.StepPackageStore/";

  DefaultStepPackageStore(this.hostname, {Requester requester}) {
    if (requester == null) {
      _requester = new Requester(new Client());
    } else {
      _requester = requester;
    }
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

  Future<Empty> registerPackage(RegisteredPackage registeredPackage) async {
    var url = "${hostname}${_pathPrefix}RegisterPackage";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(registeredPackage.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return Empty.fromJson(value);
  }

  Future<RegisteredPackage> getPackage(
      StepPackageRequest stepPackageRequest) async {
    var url = "${hostname}${_pathPrefix}GetPackage";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(stepPackageRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return RegisteredPackage.fromJson(value);
  }

  Future<RegisteredPackage> getPackageById(
      StepPackageIdRequest stepPackageIdRequest) async {
    var url = "${hostname}${_pathPrefix}GetPackageById";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(stepPackageIdRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return RegisteredPackage.fromJson(value);
  }

  Future<GetSingleStepResponse> getStep(
      GetSingleStepRequest getSingleStepRequest) async {
    var url = "${hostname}${_pathPrefix}GetStep";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(getSingleStepRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return GetSingleStepResponse.fromJson(value);
  }

  Future<SearchStepsResponse> searchStepsForLibrary(
      SearchStepsRequest searchStepsRequest) async {
    var url = "${hostname}${_pathPrefix}SearchStepsForLibrary";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(searchStepsRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return SearchStepsResponse.fromJson(value);
  }

  Future<GetAllStepsResponse> getAllStepsForLibrary(
      GetAllStepsRequest getAllStepsRequest) async {
    var url = "${hostname}${_pathPrefix}GetAllStepsForLibrary";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(getAllStepsRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return GetAllStepsResponse.fromJson(value);
  }

  Future<GetStepsForPackageResponse> getPackageSteps(
      GetStepsForPackageRequest getStepsForPackageRequest) async {
    var url = "${hostname}${_pathPrefix}GetPackageSteps";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(getStepsForPackageRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return GetStepsForPackageResponse.fromJson(value);
  }

  Future<GetAllPackagesInfoResponse> getAllPackageNamesForLibrary(
      GetAllPackagesInfoRequest getAllPackagesInfoRequest) async {
    var url = "${hostname}${_pathPrefix}GetAllPackageNamesForLibrary";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(getAllPackagesInfoRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return GetAllPackagesInfoResponse.fromJson(value);
  }

  Future<GetPackageInfoResponse> getPackageInfoForLibrary(
      GetPackageInfoRequest getPackageInfoRequest) async {
    var url = "${hostname}${_pathPrefix}GetPackageInfoForLibrary";
    var uri = Uri.parse(url);
    var request = new Request('POST', uri);
    request.headers['Content-Type'] = 'application/json';
    request.body = json.encode(getPackageInfoRequest.toJson());
    var response = await _requester.send(request);
    if (response.statusCode != 200) {
      throw twirpException(response);
    }
    var value = json.decode(response.body);
    return GetPackageInfoResponse.fromJson(value);
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
