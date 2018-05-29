package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var mongoClient *mongo.Client

// 建立与 MongoDB 的连接，必须使用 auth 模式
func connectMongo(url string) {
	var err error
	opts := &mongo.ClientOptions{}
	opts = opts.AuthMechanism(cfg.MongoAuthMechanism).Username(cfg.MongoUsername).
		Password(cfg.MongoPassword).AuthSource(cfg.MongoAuthSource)
	mongoClient, err = mongo.Connect(context.Background(), url, opts)
	if err != nil {
		log.Errorf("[Error] Connect mongo failed: %s", err.Error())
	}
}

// 插入文档到指定 database, collection
func insertMongoDocument(database string, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	return mongoClient.Database(database).Collection(collection).InsertOne(context.Background(), document)
}
