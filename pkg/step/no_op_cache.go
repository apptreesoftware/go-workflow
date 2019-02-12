package step

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type NoOpCache struct {
}

func (NoOpCache) Push(ctx context.Context, in *CachePushRequest, opts ...grpc.CallOption) (*CachePushResponse, error) {
	return &CachePushResponse{}, nil
}

func (NoOpCache) Pull(ctx context.Context, in *CachePullRequest, opts ...grpc.CallOption) (*CachePullResponse, error) {
	return &CachePullResponse{}, nil
}
