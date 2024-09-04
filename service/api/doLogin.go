package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

// doLogin handles the user login request
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Parse the incoming user data from the request body
	var loginUser User
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Failed to parse user data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Attempt to create or verify the user in the database
	dbUser, err := rt.db.CreateUser(loginUser.CovertForDatabase())
	if err != nil {
		http.Error(w, "Database operation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Populate the user object with the data from the database
	loginUser.ConvertForApplication(dbUser)

	// Prepare the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(loginUser); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
