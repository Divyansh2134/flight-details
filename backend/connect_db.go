package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo() (*mongo.Client, error) {
	if env.MongoDbUrl == "" {
		env.MongoDbUrl = "mongodb://localhost:27017"
	}
	clientOptions := options.Client().ApplyURI(env.MongoDbUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		Log.Error(err)
	}

	return client, nil
}
