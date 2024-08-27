package api

import (
	"encoding/json"
	"net/http"

	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var requestUser User
	var profile Profile
	token := getToken(r.Header.Get("Authorization"))
	requestUser.Id = token
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requestUser.FromDatabase(dbrequestuser)
	username := ps.ByName("username")
	userid, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile.RequestId = token
	profile.Id = userid
	profile.Username = username
	followersCount, err := rt.db.GetFollowersCount(userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.FollowersCount = followersCount
	followingCount, err := rt.db.GetFollowingsCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.FollowingCount = followingCount
	photoCount, err := rt.db.GetPhotosCount(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.PhotoCount = photoCount
	profile.BanStatus, err = rt.db.GetBanStatus(requestUser.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.FollowStatus, err = rt.db.GetFollowStatus(requestUser.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	profile.CheckIfBanned, err = rt.db.CheckIfBanned(requestUser.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(profile)
}
