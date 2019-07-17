// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: step.proto
// </auto-generated>
// Original file comments:
// Defines the API for a step using the SDK to talk to an engine.
//
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Core {
  public static partial class EngineStepAPI
  {
    static readonly string __ServiceName = "core.EngineStepAPI";

    static readonly grpc::Marshaller<global::Core.EmptyMessage> __Marshaller_core_EmptyMessage = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.EmptyMessage.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.CachePushRequest> __Marshaller_core_CachePushRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.CachePushRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.CachePushResponse> __Marshaller_core_CachePushResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.CachePushResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.CachePullRequest> __Marshaller_core_CachePullRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.CachePullRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.CachePullResponse> __Marshaller_core_CachePullResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.CachePullResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.CacheSearchRequest> __Marshaller_core_CacheSearchRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.CacheSearchRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.CacheSearchResponse> __Marshaller_core_CacheSearchResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.CacheSearchResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.StepQueueWorkflowRequest> __Marshaller_core_StepQueueWorkflowRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.StepQueueWorkflowRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.StepQueueWorkflowResponse> __Marshaller_core_StepQueueWorkflowResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.StepQueueWorkflowResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.GetWorkflowUrlRequest> __Marshaller_core_GetWorkflowUrlRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.GetWorkflowUrlRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.GetWorkflowUrlResponse> __Marshaller_core_GetWorkflowUrlResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.GetWorkflowUrlResponse.Parser.ParseFrom);

    static readonly grpc::Method<global::Core.EmptyMessage, global::Core.EmptyMessage> __Method_Ping = new grpc::Method<global::Core.EmptyMessage, global::Core.EmptyMessage>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Ping",
        __Marshaller_core_EmptyMessage,
        __Marshaller_core_EmptyMessage);

    static readonly grpc::Method<global::Core.CachePushRequest, global::Core.CachePushResponse> __Method_CachePush = new grpc::Method<global::Core.CachePushRequest, global::Core.CachePushResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CachePush",
        __Marshaller_core_CachePushRequest,
        __Marshaller_core_CachePushResponse);

    static readonly grpc::Method<global::Core.CachePullRequest, global::Core.CachePullResponse> __Method_CachePull = new grpc::Method<global::Core.CachePullRequest, global::Core.CachePullResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CachePull",
        __Marshaller_core_CachePullRequest,
        __Marshaller_core_CachePullResponse);

    static readonly grpc::Method<global::Core.CacheSearchRequest, global::Core.CacheSearchResponse> __Method_CacheSearch = new grpc::Method<global::Core.CacheSearchRequest, global::Core.CacheSearchResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CacheSearch",
        __Marshaller_core_CacheSearchRequest,
        __Marshaller_core_CacheSearchResponse);

    static readonly grpc::Method<global::Core.StepQueueWorkflowRequest, global::Core.StepQueueWorkflowResponse> __Method_QueueWorkflow = new grpc::Method<global::Core.StepQueueWorkflowRequest, global::Core.StepQueueWorkflowResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "QueueWorkflow",
        __Marshaller_core_StepQueueWorkflowRequest,
        __Marshaller_core_StepQueueWorkflowResponse);

    static readonly grpc::Method<global::Core.GetWorkflowUrlRequest, global::Core.GetWorkflowUrlResponse> __Method_GetWorkflowUrl = new grpc::Method<global::Core.GetWorkflowUrlRequest, global::Core.GetWorkflowUrlResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetWorkflowUrl",
        __Marshaller_core_GetWorkflowUrlRequest,
        __Marshaller_core_GetWorkflowUrlResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Core.StepReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of EngineStepAPI</summary>
    public abstract partial class EngineStepAPIBase
    {
      public virtual global::System.Threading.Tasks.Task<global::Core.EmptyMessage> Ping(global::Core.EmptyMessage request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Core.CachePushResponse> CachePush(global::Core.CachePushRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Core.CachePullResponse> CachePull(global::Core.CachePullRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Core.CacheSearchResponse> CacheSearch(global::Core.CacheSearchRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Core.StepQueueWorkflowResponse> QueueWorkflow(global::Core.StepQueueWorkflowRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Core.GetWorkflowUrlResponse> GetWorkflowUrl(global::Core.GetWorkflowUrlRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for EngineStepAPI</summary>
    public partial class EngineStepAPIClient : grpc::ClientBase<EngineStepAPIClient>
    {
      /// <summary>Creates a new client for EngineStepAPI</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public EngineStepAPIClient(grpc::Channel channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for EngineStepAPI that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public EngineStepAPIClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected EngineStepAPIClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected EngineStepAPIClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::Core.EmptyMessage Ping(global::Core.EmptyMessage request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Ping(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Core.EmptyMessage Ping(global::Core.EmptyMessage request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Ping, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Core.EmptyMessage> PingAsync(global::Core.EmptyMessage request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return PingAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Core.EmptyMessage> PingAsync(global::Core.EmptyMessage request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Ping, null, options, request);
      }
      public virtual global::Core.CachePushResponse CachePush(global::Core.CachePushRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CachePush(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Core.CachePushResponse CachePush(global::Core.CachePushRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CachePush, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Core.CachePushResponse> CachePushAsync(global::Core.CachePushRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CachePushAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Core.CachePushResponse> CachePushAsync(global::Core.CachePushRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CachePush, null, options, request);
      }
      public virtual global::Core.CachePullResponse CachePull(global::Core.CachePullRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CachePull(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Core.CachePullResponse CachePull(global::Core.CachePullRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CachePull, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Core.CachePullResponse> CachePullAsync(global::Core.CachePullRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CachePullAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Core.CachePullResponse> CachePullAsync(global::Core.CachePullRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CachePull, null, options, request);
      }
      public virtual global::Core.CacheSearchResponse CacheSearch(global::Core.CacheSearchRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CacheSearch(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Core.CacheSearchResponse CacheSearch(global::Core.CacheSearchRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CacheSearch, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Core.CacheSearchResponse> CacheSearchAsync(global::Core.CacheSearchRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CacheSearchAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Core.CacheSearchResponse> CacheSearchAsync(global::Core.CacheSearchRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CacheSearch, null, options, request);
      }
      public virtual global::Core.StepQueueWorkflowResponse QueueWorkflow(global::Core.StepQueueWorkflowRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return QueueWorkflow(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Core.StepQueueWorkflowResponse QueueWorkflow(global::Core.StepQueueWorkflowRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_QueueWorkflow, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Core.StepQueueWorkflowResponse> QueueWorkflowAsync(global::Core.StepQueueWorkflowRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return QueueWorkflowAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Core.StepQueueWorkflowResponse> QueueWorkflowAsync(global::Core.StepQueueWorkflowRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_QueueWorkflow, null, options, request);
      }
      public virtual global::Core.GetWorkflowUrlResponse GetWorkflowUrl(global::Core.GetWorkflowUrlRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetWorkflowUrl(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Core.GetWorkflowUrlResponse GetWorkflowUrl(global::Core.GetWorkflowUrlRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetWorkflowUrl, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Core.GetWorkflowUrlResponse> GetWorkflowUrlAsync(global::Core.GetWorkflowUrlRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetWorkflowUrlAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Core.GetWorkflowUrlResponse> GetWorkflowUrlAsync(global::Core.GetWorkflowUrlRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetWorkflowUrl, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override EngineStepAPIClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new EngineStepAPIClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(EngineStepAPIBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_Ping, serviceImpl.Ping)
          .AddMethod(__Method_CachePush, serviceImpl.CachePush)
          .AddMethod(__Method_CachePull, serviceImpl.CachePull)
          .AddMethod(__Method_CacheSearch, serviceImpl.CacheSearch)
          .AddMethod(__Method_QueueWorkflow, serviceImpl.QueueWorkflow)
          .AddMethod(__Method_GetWorkflowUrl, serviceImpl.GetWorkflowUrl).Build();
    }

    /// <summary>Register service method implementations with a service binder. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static void BindService(grpc::ServiceBinderBase serviceBinder, EngineStepAPIBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_Ping, serviceImpl.Ping);
      serviceBinder.AddMethod(__Method_CachePush, serviceImpl.CachePush);
      serviceBinder.AddMethod(__Method_CachePull, serviceImpl.CachePull);
      serviceBinder.AddMethod(__Method_CacheSearch, serviceImpl.CacheSearch);
      serviceBinder.AddMethod(__Method_QueueWorkflow, serviceImpl.QueueWorkflow);
      serviceBinder.AddMethod(__Method_GetWorkflowUrl, serviceImpl.GetWorkflowUrl);
    }

  }
  public static partial class StepHost
  {
    static readonly string __ServiceName = "core.StepHost";

    static readonly grpc::Marshaller<global::Core.RunStepRequest> __Marshaller_core_RunStepRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.RunStepRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.StepOutput> __Marshaller_core_StepOutput = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.StepOutput.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.EmptyMessage> __Marshaller_core_EmptyMessage = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.EmptyMessage.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Core.Package> __Marshaller_core_Package = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Core.Package.Parser.ParseFrom);

    static readonly grpc::Method<global::Core.RunStepRequest, global::Core.StepOutput> __Method_RunStep = new grpc::Method<global::Core.RunStepRequest, global::Core.StepOutput>(
        grpc::MethodType.Unary,
        __ServiceName,
        "RunStep",
        __Marshaller_core_RunStepRequest,
        __Marshaller_core_StepOutput);

    static readonly grpc::Method<global::Core.EmptyMessage, global::Core.Package> __Method_GetPackageInfo = new grpc::Method<global::Core.EmptyMessage, global::Core.Package>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetPackageInfo",
        __Marshaller_core_EmptyMessage,
        __Marshaller_core_Package);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Core.StepReflection.Descriptor.Services[1]; }
    }

    /// <summary>Base class for server-side implementations of StepHost</summary>
    public abstract partial class StepHostBase
    {
      public virtual global::System.Threading.Tasks.Task<global::Core.StepOutput> RunStep(global::Core.RunStepRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Core.Package> GetPackageInfo(global::Core.EmptyMessage request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for StepHost</summary>
    public partial class StepHostClient : grpc::ClientBase<StepHostClient>
    {
      /// <summary>Creates a new client for StepHost</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public StepHostClient(grpc::Channel channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for StepHost that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public StepHostClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected StepHostClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected StepHostClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::Core.StepOutput RunStep(global::Core.RunStepRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RunStep(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Core.StepOutput RunStep(global::Core.RunStepRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_RunStep, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Core.StepOutput> RunStepAsync(global::Core.RunStepRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RunStepAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Core.StepOutput> RunStepAsync(global::Core.RunStepRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_RunStep, null, options, request);
      }
      public virtual global::Core.Package GetPackageInfo(global::Core.EmptyMessage request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetPackageInfo(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Core.Package GetPackageInfo(global::Core.EmptyMessage request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetPackageInfo, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Core.Package> GetPackageInfoAsync(global::Core.EmptyMessage request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetPackageInfoAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Core.Package> GetPackageInfoAsync(global::Core.EmptyMessage request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetPackageInfo, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override StepHostClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new StepHostClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(StepHostBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_RunStep, serviceImpl.RunStep)
          .AddMethod(__Method_GetPackageInfo, serviceImpl.GetPackageInfo).Build();
    }

    /// <summary>Register service method implementations with a service binder. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static void BindService(grpc::ServiceBinderBase serviceBinder, StepHostBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_RunStep, serviceImpl.RunStep);
      serviceBinder.AddMethod(__Method_GetPackageInfo, serviceImpl.GetPackageInfo);
    }

  }
}
#endregion
