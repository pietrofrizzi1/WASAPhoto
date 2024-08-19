package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Comment struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var Comments = [][]Comment{
	{
		Comment{Username: "user1", Text: "First comment"},
		Comment{Username: "user2", Text: "Second comment"},
		Comment{Username: "user3", Text: "Third comment"},
	},
}

func (rt *_router) CommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the comment from query parameters
	commentText := r.URL.Query().Get("comment")

	// Create a new Comment instance
	newComment := Comment{
		Username: "Anonymous", // Assuming no username is provided
		Text:     commentText,
	}

	// Append the new comment to Comments
	Comments[0] = append(Comments[0], newComment)

	// Encode the newly created comment into JSON format and send it as the response
	json.NewEncoder(w).Encode(newComment)
}
