package cache

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
)

var (
	MongoDBPool *mongo.Client
	MongoHost   string
	MongoPort   string
)

func loadMongoDBData(file *ini.File) {
	MongoHost = file.Section("mongo").Key("MongoHost").String()
	MongoPort = file.Section("mongo").Key("MongoPort").String()
}

func newMongoDBPool() *mongo.Client {
	// 建立连接
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://"+MongoHost+":"+MongoPort).SetMaxPoolSize(20))
	if err != nil {
		fmt.Printf("mongo connect err:%v\n", err)
	}
	return client
}
