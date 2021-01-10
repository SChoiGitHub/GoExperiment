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

func getClient(){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Connect(ctx)
	defer cancel();

	disconnect = func() {
		if err = client.Disconnect()
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
	messages := client.Database("testing").Collection("message")
	res, _ := collection.InsertOne(ctx, bson.M{"value": message})

	fmt.Fprint(w, "Saved: ", message, " at ", res.InsertedID)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/echo/{message}", EchoHandler)
	router.HandleFunc("/save/{message}", SaveHandler)

	http.ListenAndServe(":5000", router)
}