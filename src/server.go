package main

import (
	"fmt"
	"net/http"
	"repository"

	"github.com/gorilla/mux"
)

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

	res := repository.SaveMessage(message)

	fmt.Fprint(w, "Saved: ", message, " at ", res.InsertedID)
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(w, "Message not found at ID.")
		}
	}()

	params := mux.Vars(r)
	id := params["id"]

	message := repository.LoadMessage(id)

	fmt.Fprint(w, "Message found: ", message.Text)
}

func main() {
	defer repository.Disconnect()

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/echo/{message}", EchoHandler)
	router.HandleFunc("/save/{message}", SaveHandler)
	router.HandleFunc("/load/{id}", LoadHandler)

	http.ListenAndServe(":5000", router)
}
