package main

import (
	"fmt"
	"net/http"
	"repository"

	"github.com/gorilla/mux"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message := params["message"]
	location := params["location"]

	repository.SaveMessage(location, message)

	fmt.Fprint(w, "Saved: ", message, " at ", location)
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(w, "Message not found at Location.")
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
	router.HandleFunc("/save/{location}/{message}", SaveHandler)
	router.HandleFunc("/load/{id}", LoadHandler)

	http.ListenAndServe(":5000", router)
}
