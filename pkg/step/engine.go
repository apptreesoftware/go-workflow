package step

import (
	"context"
	. "github.com/apptreesoftware/go-workflow/pkg/core"
	jsoniter "github.com/json-iterator/go"
	"github.com/mongodb/mongo-go-driver/bson"
	"google.golang.org/grpc"
)

type Engine struct {
	client      EngineStepAPIClient
	environment *Environment
}

func GetEngine(environment *Environment) Engine {
	connectionString := environment.GetCacheHost()
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
	return Engine{client: client, environment: environment}
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
			Environment: c.environment,
		})
	return err
}

func (c *Engine) PullRecord(id string, record interface{}, cacheName string) (found bool, metadata map[string]interface{}, err error) {
	resp, pullErr := c.client.CachePull(context.Background(), &CachePullRequest{
		Id:          id,
		CacheName:   cacheName,
		Environment: c.environment,
	})
	if pullErr != nil {
		err = pullErr
		return
	}
	if resp.NotFound {
		return
	}
	found = true
	err = bson.Unmarshal(resp.Metadata, &metadata)
	if err != nil {
		return
	}
	err = bson.Unmarshal(resp.Record, record)
	return
}

func (c *Engine) AddToQueue(workflow string, record interface{}) (err error) {
	var recordBytes []byte
	if record != nil {
		if str, ok := record.(string); ok {
			recordBytes = []byte(str)
		} else {
			recordBytes, err = jsoniter.Marshal(record)
			if err != nil {
				return err
			}
		}
	}

	_, addErr := c.client.QueueWorkflow(context.Background(), &StepQueueWorkflowRequest{
		Environment: c.environment,
		Workflow:    workflow,
		Input:       recordBytes,
	})
	err = addErr
	return
}

func (c *Engine) Find(filter interface{}, cacheName string, limit int64) ([]*RawRecord, error) {
	filterBytes, err := bson.Marshal(&filter)
	if err != nil {
		return nil, err
	}
	searchRequest := &CacheSearchRequest{
		CacheName:    cacheName,
		Environment:  c.environment,
		SearchFilter: filterBytes,
		Limit:        limit,
	}
	resp, err := c.client.CacheSearch(context.Background(), searchRequest)
	if err != nil {
		return nil, err
	}
	return resp.Records, nil
}
