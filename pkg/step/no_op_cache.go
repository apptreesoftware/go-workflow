package step

import (
	. "github.com/apptreesoftware/go-workflow/pkg/core"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	
)

type NoOpEngine struct {
}

func (NoOpEngine) QueueWorkflow(ctx context.Context, in *StepQueueWorkflowRequest, opts ...grpc.CallOption) (*StepQueueWorkflowResponse, error) {
	return &StepQueueWorkflowResponse{}, nil
}

func (NoOpEngine) CachePush(ctx context.Context, in *CachePushRequest, opts ...grpc.CallOption) (*CachePushResponse, error) {
	return &CachePushResponse{}, nil
}

func (NoOpEngine) CachePull(ctx context.Context, in *CachePullRequest, opts ...grpc.CallOption) (*CachePullResponse, error) {
	return &CachePullResponse{}, nil
}
