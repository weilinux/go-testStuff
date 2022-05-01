package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("hello test"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
