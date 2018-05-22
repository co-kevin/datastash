package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

var mongoClient *mongo.Client

func init() {
	connectMongo("mongodb://localhost:27017")
}

func connectMongo(url string) {
	var err error
	mongoClient, err = mongo.Connect(context.Background(), url, nil)
	if err != nil {
		log.Println("[Error] connect mongo: " + err.Error())
	}
}

func insertMongoDocument(database string, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	return mongoClient.Database(database).Collection(collection).InsertOne(context.Background(), document)
}
