package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

//RPCServer是我们的RPC服务器的类型。可以使用将其作为接收器的方法
// over RPC, as long as they are exported.
type RPCServer struct{}

// RPCPayload是我们从RPC接收的数据类型
type RPCPayload struct {
	Name string
	Data string
}

//LogInfo将payload写入Mongo
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	//resp是发送回RPC调用方的消息
	*resp = "Processed payload via RPC:" + payload.Name
	return nil
}
