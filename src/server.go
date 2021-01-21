package main

import (
	"fmt"
	"net/http"
	"repository"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Text string
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

	res, err := repository.MessageCollection.InsertOne(repository.Context, bson.M{"Text": message})

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

	err = repository.MessageCollection.FindOne(repository.Context, bson.M{"_id": objId}).Decode(&message)

	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, "Message found: ", message.Text)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/echo/{message}", EchoHandler)
	router.HandleFunc("/save/{message}", SaveHandler)
	router.HandleFunc("/load/{id}", LoadHandler)

	http.ListenAndServe(":5000", router)
}
