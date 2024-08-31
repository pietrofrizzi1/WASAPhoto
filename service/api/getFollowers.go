package api

import (
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var follow Follow
	token := getToken(r.Header.Get("Authorization"))
	user.Username = ps.ByName("username")
	dbuser, err := rt.db.CheckUserByUsername(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	dbfollow, err := rt.db.GetFollowers(user.ToDatabase(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	follow.FollowFromDatabase(dbfollow)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(follow)
}
