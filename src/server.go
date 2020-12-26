package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello!")
}

func EchoHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	fmt.Fprint(w, "Echo: ", params["message"])
}

func SaveHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	message := params["message"];
	fmt.Fprint(w, "Echo: ", message)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/echo/{message}", EchoHandler)
	router.HandleFunc("/save/{message}", SaveHandler)

	http.ListenAndServe(":5000", router)
}