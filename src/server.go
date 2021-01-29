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
	defer recoverIfMessageNotFound(w)

	params := mux.Vars(r)
	location := params["location"]

	message := repository.LoadMessage(location)

	fmt.Fprint(w, "Message found: ", message.Text)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	defer recoverIfMessageNotFound(w)

	params := mux.Vars(r)
	location := params["location"]

	repository.DeleteMessage(location)

	fmt.Fprint(w, "Deleted message at: ", location)
}

func recoverIfMessageNotFound(w http.ResponseWriter) {
	if r := recover(); r != nil {
		fmt.Fprint(w, "Message not found at location.")
	}
}

func main() {
	defer repository.Disconnect()

	router := mux.NewRouter()
	router.HandleFunc("/save/{location}/{message}", SaveHandler)
	router.HandleFunc("/load/{location}", LoadHandler)
	router.HandleFunc("/delete/{location}", DeleteHandler)

	http.ListenAndServe(":5000", router)
}
