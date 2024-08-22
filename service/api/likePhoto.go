package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like
	likeid, err := strconv.ParseUint(ps.ByName("likeid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photoid, err := strconv.ParseUint(ps.ByName("photoid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := ps.ByName("username")
	userid, err := rt.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := getToken(r.Header.Get("Authorization"))

	Like.Id = likeid
	Like.PhotoId = photoid
	Like.PhotoOwner = token
	Like.UserId = userid

	mipiaze, err := rt.db.SetLike(Like.LikeToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Like.LikeFromDatabase(mipiaze)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
