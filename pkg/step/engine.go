package step

import (
	"context"
	. "github.com/apptreesoftware/go-workflow/pkg/core"
	"github.com/json-iterator/go"
	"github.com/mongodb/mongo-go-driver/bson"
	"google.golang.org/grpc"
	"os"
)

type Engine struct {
	client EngineStepAPIClient
}


func GetEngine() Engine {
	connectionString := os.Getenv("WORKFLOW_CACHE_CONNECTION")
	println(connectionString)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	if connectionString == "" {
		return Engine{client: &NoOpEngine{}}
	}
	conn, err := grpc.Dial(connectionString, opts...)
	if err != nil {
		panic("Could not connect to engine")
	}
	client := NewEngineStepAPIClient(conn)
	return Engine{client: client}
}

func (c *Engine) PutRecord(id string, record interface{}, metadata map[string]interface{}, cacheName string) error {
	recordBytes, err := bson.Marshal(record)
	if err != nil {
		return err
	}
	metaBytes, err := bson.Marshal(metadata)
	if err != nil {
		return err
	}
	_, err = c.client.CachePush(context.Background(),
		&CachePushRequest{
			Id:          id,
			Metadata:    metaBytes,
			Record:      recordBytes,
			CacheName:   cacheName,
			Environment: GetEnvironment(),
		})
	return err
}

func (c *Engine) PullRecord(id string, record interface{}, cacheName string) (metadata map[string]interface{}, err error) {
	environment := GetEnvironment()
	resp, err := c.client.CachePull(context.Background(), &CachePullRequest{
		Id:          id,
		CacheName:   cacheName,
		Environment: environment,
	})
	err = bson.Unmarshal(resp.Metadata, &metadata)
	if err != nil {
		return
	}
	err = bson.Unmarshal(resp.Record, &record)
	return
}

func (c *Engine) AddToQueue(workflow string, record interface{}) (err error) {
	environment := GetEnvironment()
	var recordBytes []byte
	if record != nil {
		recordBytes, err = jsoniter.Marshal(record)
		if err != nil {
			return err
		}
	}

	_, addErr := c.client.QueueWorkflow(context.Background(), &StepQueueWorkflowRequest{
		Environment:          environment,
		Workflow:             workflow,
		Input:                recordBytes,
	})
	err = addErr
	return
}