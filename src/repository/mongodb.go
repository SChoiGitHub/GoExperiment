package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type message struct {
	location string
	Text     string
}

var ctx = context.Background()

func initializeClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	return client
}

var Client = initializeClient()
var messageCollection = Client.Database("testing").Collection("message")

func Disconnect() {
	err := Client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}

func SaveMessage(location string, message string) {
	_, err := messageCollection.InsertOne(ctx, bson.M{
		"Location": location,
		"Text":     message,
	})

	if err != nil {
		panic(err)
	}
}

func LoadMessage(idAsString string) message {
	var message message
	objId, err := primitive.ObjectIDFromHex(idAsString)
	err = messageCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&message)

	if err != nil {
		panic(err)
	}

	return message
}
