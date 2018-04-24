package main

import (
	"goblog/blog"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/post", blog.GetPost).Methods("GET")
	router.HandleFunc("/post/{id}", blog.GetPost).Methods("GET")
	router.HandleFunc("/post/{id}", blog.CreatePost).Methods("POST")
	router.HandleFunc("/post/{id}", blog.DeletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
