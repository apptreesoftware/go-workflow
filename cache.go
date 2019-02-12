package go_workflow

import (
	"context"
	. "github.com/apptreesoftware/go-workflow/pkg/cache"
	"github.com/mongodb/mongo-go-driver/bson"
	"google.golang.org/grpc"
	"os"
)

type Cache struct {
	client CacheClient
}

func GetCache() Cache {
	connectionString := os.Getenv("WORKFLOW_CACHE_CONNECTION")
	println(connectionString)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	if connectionString == "" {
		return Cache{client: &NoOpCache{}}
	}
	conn, err := grpc.Dial(connectionString, opts...)
	if err != nil {
		panic("Could not connect to cache")
	}
	client := NewCacheClient(conn)
	return Cache{client: client}
}

func (c *Cache) PutRecord(id string, record interface{}, metadata map[string]interface{}, cacheName string) error {
	recordBytes, err := bson.Marshal(record)
	if err != nil {
		return err
	}
	metaBytes, err := bson.Marshal(metadata)
	if err != nil {
		return err
	}
	_, err = c.client.Push(context.Background(),
		&CachePushRequest{
			Id:          id,
			Metadata:    metaBytes,
			Record:      recordBytes,
			CacheName:   cacheName,
			Environment: GetEnvironment(),
		})
	return err
}

func GetEnvironment() *Environment {
	runId := os.Getenv("RUN_ID")
	appId := os.Getenv("APP")
	workflowId := os.Getenv("WORKFLOW_ID")
	return &Environment{
		App:        appId,
		WorkflowId: workflowId,
		RunId:      runId,
	}
}

func (c *Cache) PullRecord(id string, record interface{}, cacheName string) (metadata map[string]interface{}, err error) {
	environment := GetEnvironment()
	resp, err := c.client.Pull(context.Background(), &CachePullRequest{
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
