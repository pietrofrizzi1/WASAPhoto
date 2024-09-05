package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

// commentPhoto handles the process of adding or updating a comment on a photo
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Decode the incoming request body to retrieve comment data
	var newComment Comment
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Extract photo ID from request parameters
	photoID, err := strconv.ParseUint(ps.ByName("singlephoto"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the username from the request parameters
	username := ps.ByName("singleusername")

	// Fetch user data from the database
	userData, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, "Error fetching user data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var user User
	user.ConvertForApplication(userData)

	// Extract comment ID from request parameters
	commentID, err := strconv.ParseUint(ps.ByName("singlecomment"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the comment object with extracted data
	newComment.Id = commentID
	token := getAuthorization(r.Header.Get("Authorization"))
	newComment.UserId = token
	newComment.PhotoId = photoID
	newComment.PhotoOwner = user.Id

	// Save the comment to the database
	savedCommentData, err := rt.db.SetComment(newComment.CommentConvertForDatabase())
	if err != nil {
		http.Error(w, "Error saving comment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Populate the comment object with saved data
	newComment.CommentConvertForApplication(savedCommentData)

	// Respond with the created comment
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newComment); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}
