package step

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/core"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"testing"
)

func TestEngine_AddToQueue(t *testing.T) {
	mockClient := mockEngineClient{}
	e := Engine{
		client:      &mockClient,
		environment: nil,
	}
	inputMap := map[string]interface{}{
		"key" : "value",
	}
	_ = e.AddToQueue("test", inputMap)
	expectedBytes, _ := jsoniter.Marshal(inputMap)
	assert.Equal(t, expectedBytes, mockClient.queueInput)

	_ = e.AddToQueue("test", "testString")
	assert.Equal(t, []byte("testString"), mockClient.queueInput)
}

type mockEngineClient struct {
	m mock.Mock
	queueInput []byte
}

func (m mockEngineClient) Ping(ctx context.Context, in *core.EmptyMessage, opts ...grpc.CallOption) (*core.EmptyMessage, error) {
	panic("implement me")
}

func (m mockEngineClient) CachePush(ctx context.Context, in *core.CachePushRequest, opts ...grpc.CallOption) (*core.CachePushResponse, error) {
	panic("implement me")
}

func (m mockEngineClient) CachePull(ctx context.Context, in *core.CachePullRequest, opts ...grpc.CallOption) (*core.CachePullResponse, error) {
	panic("implement me")
}

func (m mockEngineClient) CacheSearch(ctx context.Context, in *core.CacheSearchRequest, opts ...grpc.CallOption) (*core.CacheSearchResponse, error) {
	panic("implement me")
}

func (m *mockEngineClient) QueueWorkflow(ctx context.Context, in *core.StepQueueWorkflowRequest, opts ...grpc.CallOption) (*core.StepQueueWorkflowResponse, error) {
	m.queueInput = in.Input
	return &core.StepQueueWorkflowResponse{}, nil
}
