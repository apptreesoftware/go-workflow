// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: step.proto
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
            "ChlTdGVwUXVldWVXb3JrZmxvd1Jlc3BvbnNlMo0CCg1FbmdpbmVTdGVwQVBJ",
            "Ei4KBFBpbmcSEi5jb3JlLkVtcHR5TWVzc2FnZRoSLmNvcmUuRW1wdHlNZXNz",
            "YWdlEjwKCUNhY2hlUHVzaBIWLmNvcmUuQ2FjaGVQdXNoUmVxdWVzdBoXLmNv",
            "cmUuQ2FjaGVQdXNoUmVzcG9uc2USPAoJQ2FjaGVQdWxsEhYuY29yZS5DYWNo",
            "ZVB1bGxSZXF1ZXN0GhcuY29yZS5DYWNoZVB1bGxSZXNwb25zZRJQCg1RdWV1",
            "ZVdvcmtmbG93Eh4uY29yZS5TdGVwUXVldWVXb3JrZmxvd1JlcXVlc3QaHy5j",
            "b3JlLlN0ZXBRdWV1ZVdvcmtmbG93UmVzcG9uc2VCMVovZ2l0aHViLmNvbS9h",
            "cHB0cmVlc29mdHdhcmUvZ28td29ya2Zsb3cvcGtnL2NvcmViBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Core.CommonReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Core.StepQueueWorkflowRequest), global::Core.StepQueueWorkflowRequest.Parser, new[]{ "Environment", "Workflow", "Input" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Core.StepQueueWorkflowResponse), global::Core.StepQueueWorkflowResponse.Parser, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class StepQueueWorkflowRequest : pb::IMessage<StepQueueWorkflowRequest> {
    private static readonly pb::MessageParser<StepQueueWorkflowRequest> _parser = new pb::MessageParser<StepQueueWorkflowRequest>(() => new StepQueueWorkflowRequest());
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
      Environment = other.environment_ != null ? other.Environment.Clone() : null;
      workflow_ = other.workflow_;
      input_ = other.input_;
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
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (environment_ != null) hash ^= Environment.GetHashCode();
      if (Workflow.Length != 0) hash ^= Workflow.GetHashCode();
      if (Input.Length != 0) hash ^= Input.GetHashCode();
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
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
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
      return true;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(StepQueueWorkflowResponse other) {
      if (other == null) {
        return;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
