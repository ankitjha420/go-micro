package main

import (
	"context"
	"log"
	"logger-service/data"
	"time"
)

type RPCServer struct {
	
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collections := client.Database("logs").Collection("logs")
	_, err := collections.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "processed payload with rpc: " + payload.Name
	return nil
}
