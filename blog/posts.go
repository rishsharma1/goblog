package blog

import (
	"net/http"
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
func GetPosts(w http.ResponseWriter, r *http.Request) {}

// GetPost returns a specfic post
func GetPost(w http.ResponseWriter, r *http.Request) {}

// CreatePost creates a post
func CreatePost(w http.ResponseWriter, r *http.Request) {}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {}
