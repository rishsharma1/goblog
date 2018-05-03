package blog

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanobox-io/golang-scribble"
)

// Post a structure that represents a blog post
type Post struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Title   string `json:"title,omitempty"`
	Type    string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

// GetPosts returns all the posts
func GetPosts(w http.ResponseWriter, r *http.Request, db *scribble.Driver) {

	posts, err := db.ReadAll("post")
	LogError(err)
	json.NewEncoder(w).Encode(posts)
}

// GetPost returns a specfic post
func GetPost(w http.ResponseWriter, r *http.Request, db *scribble.Driver) {

	params := mux.Vars(r)
	posts, err := db.ReadAll("post")
	LogError(err)

	// go through the posts and return the one that matches the request
	for _, p := range posts {
		post := Post{}
		err := json.Unmarshal([]byte(p), &post)
		LogError(err)
		if post.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// return an empty post if nothing can be found
	json.NewEncoder(w).Encode(&Post{})
}

// CreatePost creates a post
func CreatePost(w http.ResponseWriter, r *http.Request, db *scribble.Driver) {

	params := mux.Vars(r)
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = params["id"]

	// create the new post
	err := db.Write("post", post.ID, post)
	LogError(err)

	// return the updated posts
	posts, err := db.ReadAll("post")
	json.NewEncoder(w).Encode(posts)
}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request, db *scribble.Driver) {

	params := mux.Vars(r)

	// delete the desired post
	err := db.Delete("post", params["id"])
	LogError(err)

	// return the updated posts
	posts, err := db.ReadAll("post")
	LogError(err)
	json.NewEncoder(w).Encode(posts)
}
