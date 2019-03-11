package step

import (
	. "github.com/apptreesoftware/go-workflow/pkg/core"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type NoOpEngine struct {
}

func (NoOpEngine) Ping(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	return &EmptyMessage{}, nil
}

func (NoOpEngine) QueueWorkflow(ctx context.Context, in *StepQueueWorkflowRequest, opts ...grpc.CallOption) (*StepQueueWorkflowResponse, error) {
	println("NoOpEngine - QueueWorkflow")
	return &StepQueueWorkflowResponse{}, nil
}

func (NoOpEngine) CachePush(ctx context.Context, in *CachePushRequest, opts ...grpc.CallOption) (*CachePushResponse, error) {
	println("NoOpEngine - CachePush")
	return &CachePushResponse{}, nil
}

func (NoOpEngine) CachePull(ctx context.Context, in *CachePullRequest, opts ...grpc.CallOption) (*CachePullResponse, error) {
	println("NoOpEngine - CachePull")
	return &CachePullResponse{}, nil
}
