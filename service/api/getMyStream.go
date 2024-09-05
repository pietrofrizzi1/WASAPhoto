package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// create user struct
	var user User
	//	create database photoList struct
	var photoList database.Stream

	// get the token from the header
	token := getAuthorization(r.Header.Get("Authorization"))
	// get the username from the url
	username := ps.ByName("singleusername")
	user.Id = token
	user.Username = username
	// get the id of the user that wants the stream
	dbuser, err := rt.db.CheckUser(user.ConvertForDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ConvertForApplication(dbuser)

	// get the stream of the user
	photos, err := rt.db.GetMyStream(user.ConvertForDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the id of the user that wants the stream
	photoList.Identifier = token
	// set the photos to the stream
	photoList.Photos = photos

	// set the header and return the stream
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photoList)
}
