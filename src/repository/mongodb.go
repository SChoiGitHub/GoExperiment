package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	Text string
}

var Context = context.Background()

func initializeClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	err = client.Connect(Context)

	if err != nil {
		panic(err)
	}

	return client
}

var Client = initializeClient()
var MessageCollection = Client.Database("testing").Collection("message")

func Disconnect() {
	err := Client.Disconnect(Context)
	if err != nil {
		panic(err)
	}
}
