package main

import (
	"goblog/blog"
	"log"
	"net/http"

	"github.com/nanobox-io/golang-scribble"

	"github.com/gorilla/mux"
)

//DATADIR contains the blog data
const DATADIR = "/tmp/blog_data"

func main() {

	// initialise
	router := mux.NewRouter()
	db, err := scribble.New(DATADIR, nil)
	blog.LogError(err)

	// get posts handler
	router.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		blog.GetPosts(w, r, db)
	}).Methods("GET")

	// get post handler
	router.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		blog.GetPost(w, r, db)
	}).Methods("GET")

	// create post handler
	router.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		blog.CreatePost(w, r, db)
	}).Methods("POST")

	// delete post handler
	router.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		blog.DeletePost(w, r, db)
	}).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8080", router))

}
