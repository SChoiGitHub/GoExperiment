package main

import (
	"context"
	"fmt"
	"hello"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	Text string
}

var ctx, cancel = context.WithCancel(context.Background())

func getClient() (*mongo.Client, func()) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	disconnect := func() {
		err = client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}

	return client, disconnect
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprint(w, "Echo: ", params["message"])
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message := params["message"]

	client, disconnect := getClient()
	defer disconnect()
	messageCollection := client.Database("testing").Collection("message")
	res, err := messageCollection.InsertOne(ctx, bson.M{"Text": message})

	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, "Saved: ", message, " at ", res.InsertedID)
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	var message Message

	params := mux.Vars(r)
	idHex := params["id"]

	objId, err := primitive.ObjectIDFromHex(idHex)

	client, disconnect := getClient()
	defer disconnect()
	messageCollection := client.Database("testing").Collection("message")
	err = messageCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&message)
	println(message.Text)

	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, "Message found: ", message.Text)
}

func main() {
	defer cancel()

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/echo/{message}", EchoHandler)
	router.HandleFunc("/save/{message}", SaveHandler)
	router.HandleFunc("/load/{id}", LoadHandler)
	router.HandleFunc("/sayHello", func(_ http.ResponseWriter, _ *http.Request) {
		hello.SayHello()
	})

	http.ListenAndServe(":5000", router)
}
