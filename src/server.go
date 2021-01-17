package main

import (
	"fmt"
	"net/http"
	"time"
	"context"
	"github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second);

func getClient() (*mongo.Client, func()) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	err = client.Connect(ctx)

	if (err != nil) {
		panic(err);
	}

	disconnect := func() {
		err = client.Disconnect(ctx);
		if (err != nil) {
			panic(err)
		}
	}

	return client, disconnect;
}

func RootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello!")
}

func EchoHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	fmt.Fprint(w, "Echo: ", params["message"])
}

func SaveHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	message := params["message"]

	client, disconnect := getClient();
	defer disconnect();
	messageCollection := client.Database("testing").Collection("message")
	res, err := messageCollection.InsertOne(ctx, bson.M{"value": message})

	if (err != nil) {
		panic(err);
	}

	fmt.Fprint(w, "Saved: ", message, " at ", res.InsertedID)
}

func main() {
	defer cancel();
	
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/echo/{message}", EchoHandler)
	router.HandleFunc("/save/{message}", SaveHandler)

	http.ListenAndServe(":5000", router)
}