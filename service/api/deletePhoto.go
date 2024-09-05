package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

// deletePhoto handles the deletion of a photo by a user
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Retrieve authorization token
	authToken := getAuthorization(r.Header.Get("Authorization"))

	// Extract the photo ID from the URL parameters
	photoID, err := strconv.ParseUint(ps.ByName("singlephoto"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid photo ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the username from the URL parameters
	username := ps.ByName("singleusername")

	// Populate the user object with the username and ID from the token
	var requestUser User
	requestUser.Username = username
	requestUser.Id = authToken

	// Verify the user in the database
	verifiedUser, err := rt.db.CheckUser(requestUser.ConvertForDatabase())
	if err != nil {
		http.Error(w, "User verification failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Populate the requestUser object with the verified user data
	requestUser.ConvertForApplication(verifiedUser)

	// Attempt to remove the photo from the database
	err = rt.db.RemovePhoto(photoID)
	if errors.Is(err, database.ErrPhotoNotFound) {
		http.Error(w, "Photo not found: "+err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("photo_id", photoID).Error("Error occurred while deleting the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with no content status indicating successful deletion
	w.WriteHeader(http.StatusNoContent)
}
