// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: step.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Core {

  /// <summary>Holder for reflection information generated from step.proto</summary>
  public static partial class StepReflection {

    #region Descriptor
    /// <summary>File descriptor for step.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static StepReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgpzdGVwLnByb3RvEgRjb3JlGgxjb21tb24ucHJvdG8iYwoYU3RlcFF1ZXVl",
            "V29ya2Zsb3dSZXF1ZXN0EiYKC2Vudmlyb25tZW50GAEgASgLMhEuY29yZS5F",
            "bnZpcm9ubWVudBIQCgh3b3JrZmxvdxgCIAEoCRINCgVpbnB1dBgDIAEoDCIb",
            "ChlTdGVwUXVldWVXb3JrZmxvd1Jlc3BvbnNlIkkKDlN0ZXBMb2dSZXF1ZXN0",
            "Eg8KB21lc3NhZ2UYASABKAkSJgoLZW52aXJvbm1lbnQYAiABKAsyES5jb3Jl",
            "LkVudmlyb25tZW50IhwKClN0ZXBPdXRwdXQSDgoGb3V0cHV0GAEgASgMMtEC",
            "Cg1FbmdpbmVTdGVwQVBJEi4KBFBpbmcSEi5jb3JlLkVtcHR5TWVzc2FnZRoS",
            "LmNvcmUuRW1wdHlNZXNzYWdlEjwKCUNhY2hlUHVzaBIWLmNvcmUuQ2FjaGVQ",
            "dXNoUmVxdWVzdBoXLmNvcmUuQ2FjaGVQdXNoUmVzcG9uc2USPAoJQ2FjaGVQ",
            "dWxsEhYuY29yZS5DYWNoZVB1bGxSZXF1ZXN0GhcuY29yZS5DYWNoZVB1bGxS",
            "ZXNwb25zZRJCCgtDYWNoZVNlYXJjaBIYLmNvcmUuQ2FjaGVTZWFyY2hSZXF1",
            "ZXN0GhkuY29yZS5DYWNoZVNlYXJjaFJlc3BvbnNlElAKDVF1ZXVlV29ya2Zs",
            "b3cSHi5jb3JlLlN0ZXBRdWV1ZVdvcmtmbG93UmVxdWVzdBofLmNvcmUuU3Rl",
            "cFF1ZXVlV29ya2Zsb3dSZXNwb25zZTJyCghTdGVwSG9zdBIxCgdSdW5TdGVw",
            "EhQuY29yZS5SdW5TdGVwUmVxdWVzdBoQLmNvcmUuU3RlcE91dHB1dBIzCg5H",
            "ZXRQYWNrYWdlSW5mbxISLmNvcmUuRW1wdHlNZXNzYWdlGg0uY29yZS5QYWNr",
            "YWdlQjFaL2dpdGh1Yi5jb20vYXBwdHJlZXNvZnR3YXJlL2dvLXdvcmtmbG93",
            "L3BrZy9jb3JlYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Core.CommonReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Core.StepQueueWorkflowRequest), global::Core.StepQueueWorkflowRequest.Parser, new[]{ "Environment", "Workflow", "Input" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Core.StepQueueWorkflowResponse), global::Core.StepQueueWorkflowResponse.Parser, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Core.StepLogRequest), global::Core.StepLogRequest.Parser, new[]{ "Message", "Environment" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Core.StepOutput), global::Core.StepOutput.Parser, new[]{ "Output" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class StepQueueWorkflowRequest : pb::IMessage<StepQueueWorkflowRequest> {
    private static readonly pb::MessageParser<StepQueueWorkflowRequest> _parser = new pb::MessageParser<StepQueueWorkflowRequest>(() => new StepQueueWorkflowRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<StepQueueWorkflowRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Core.StepReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepQueueWorkflowRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepQueueWorkflowRequest(StepQueueWorkflowRequest other) : this() {
      environment_ = other.environment_ != null ? other.environment_.Clone() : null;
      workflow_ = other.workflow_;
      input_ = other.input_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepQueueWorkflowRequest Clone() {
      return new StepQueueWorkflowRequest(this);
    }

    /// <summary>Field number for the "environment" field.</summary>
    public const int EnvironmentFieldNumber = 1;
    private global::Core.Environment environment_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Core.Environment Environment {
      get { return environment_; }
      set {
        environment_ = value;
      }
    }

    /// <summary>Field number for the "workflow" field.</summary>
    public const int WorkflowFieldNumber = 2;
    private string workflow_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Workflow {
      get { return workflow_; }
      set {
        workflow_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "input" field.</summary>
    public const int InputFieldNumber = 3;
    private pb::ByteString input_ = pb::ByteString.Empty;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pb::ByteString Input {
      get { return input_; }
      set {
        input_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as StepQueueWorkflowRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(StepQueueWorkflowRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Environment, other.Environment)) return false;
      if (Workflow != other.Workflow) return false;
      if (Input != other.Input) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (environment_ != null) hash ^= Environment.GetHashCode();
      if (Workflow.Length != 0) hash ^= Workflow.GetHashCode();
      if (Input.Length != 0) hash ^= Input.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (environment_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Environment);
      }
      if (Workflow.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Workflow);
      }
      if (Input.Length != 0) {
        output.WriteRawTag(26);
        output.WriteBytes(Input);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (environment_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Environment);
      }
      if (Workflow.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Workflow);
      }
      if (Input.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeBytesSize(Input);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(StepQueueWorkflowRequest other) {
      if (other == null) {
        return;
      }
      if (other.environment_ != null) {
        if (environment_ == null) {
          environment_ = new global::Core.Environment();
        }
        Environment.MergeFrom(other.Environment);
      }
      if (other.Workflow.Length != 0) {
        Workflow = other.Workflow;
      }
      if (other.Input.Length != 0) {
        Input = other.Input;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (environment_ == null) {
              environment_ = new global::Core.Environment();
            }
            input.ReadMessage(environment_);
            break;
          }
          case 18: {
            Workflow = input.ReadString();
            break;
          }
          case 26: {
            Input = input.ReadBytes();
            break;
          }
        }
      }
    }

  }

  public sealed partial class StepQueueWorkflowResponse : pb::IMessage<StepQueueWorkflowResponse> {
    private static readonly pb::MessageParser<StepQueueWorkflowResponse> _parser = new pb::MessageParser<StepQueueWorkflowResponse>(() => new StepQueueWorkflowResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<StepQueueWorkflowResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Core.StepReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepQueueWorkflowResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepQueueWorkflowResponse(StepQueueWorkflowResponse other) : this() {
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepQueueWorkflowResponse Clone() {
      return new StepQueueWorkflowResponse(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as StepQueueWorkflowResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(StepQueueWorkflowResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(StepQueueWorkflowResponse other) {
      if (other == null) {
        return;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
        }
      }
    }

  }

  public sealed partial class StepLogRequest : pb::IMessage<StepLogRequest> {
    private static readonly pb::MessageParser<StepLogRequest> _parser = new pb::MessageParser<StepLogRequest>(() => new StepLogRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<StepLogRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Core.StepReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepLogRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepLogRequest(StepLogRequest other) : this() {
      message_ = other.message_;
      environment_ = other.environment_ != null ? other.environment_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepLogRequest Clone() {
      return new StepLogRequest(this);
    }

    /// <summary>Field number for the "message" field.</summary>
    public const int MessageFieldNumber = 1;
    private string message_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Message {
      get { return message_; }
      set {
        message_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "environment" field.</summary>
    public const int EnvironmentFieldNumber = 2;
    private global::Core.Environment environment_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Core.Environment Environment {
      get { return environment_; }
      set {
        environment_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as StepLogRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(StepLogRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Message != other.Message) return false;
      if (!object.Equals(Environment, other.Environment)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Message.Length != 0) hash ^= Message.GetHashCode();
      if (environment_ != null) hash ^= Environment.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Message.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Message);
      }
      if (environment_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Environment);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Message.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Message);
      }
      if (environment_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Environment);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(StepLogRequest other) {
      if (other == null) {
        return;
      }
      if (other.Message.Length != 0) {
        Message = other.Message;
      }
      if (other.environment_ != null) {
        if (environment_ == null) {
          environment_ = new global::Core.Environment();
        }
        Environment.MergeFrom(other.Environment);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Message = input.ReadString();
            break;
          }
          case 18: {
            if (environment_ == null) {
              environment_ = new global::Core.Environment();
            }
            input.ReadMessage(environment_);
            break;
          }
        }
      }
    }

  }

  public sealed partial class StepOutput : pb::IMessage<StepOutput> {
    private static readonly pb::MessageParser<StepOutput> _parser = new pb::MessageParser<StepOutput>(() => new StepOutput());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<StepOutput> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Core.StepReflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepOutput() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepOutput(StepOutput other) : this() {
      output_ = other.output_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public StepOutput Clone() {
      return new StepOutput(this);
    }

    /// <summary>Field number for the "output" field.</summary>
    public const int OutputFieldNumber = 1;
    private pb::ByteString output_ = pb::ByteString.Empty;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pb::ByteString Output {
      get { return output_; }
      set {
        output_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as StepOutput);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(StepOutput other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Output != other.Output) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Output.Length != 0) hash ^= Output.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Output.Length != 0) {
        output.WriteRawTag(10);
        output.WriteBytes(Output);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Output.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeBytesSize(Output);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(StepOutput other) {
      if (other == null) {
        return;
      }
      if (other.Output.Length != 0) {
        Output = other.Output;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Output = input.ReadBytes();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
