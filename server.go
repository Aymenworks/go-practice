package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8000"

	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Println(http.ListenAndServe(port, router))
}
