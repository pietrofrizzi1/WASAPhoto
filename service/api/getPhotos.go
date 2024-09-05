package api

import (
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var requestUser User
	var photoList database.Photos
	token := getAuthorization(r.Header.Get("Authorization"))
	requestUser.Id = token
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ConvertForDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requestUser.ConvertForApplication(dbrequestuser)
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ConvertForApplication(dbuser)
	photos, err := rt.db.GetPhotos(user.ConvertForDatabase(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photoList.RequestUser = requestUser.Id
	photoList.Identifier = token
	photoList.Photos = photos
	w.Header().Set("Content-Type", "image/*")
	_ = json.NewEncoder(w).Encode(photoList)
}
