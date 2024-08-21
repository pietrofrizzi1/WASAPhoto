package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow
	id, err := strconv.ParseUint(ps.ByName("followid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datafollow, err := rt.db.GetFollowingId(id)
	if err != nil {
		if errors.Is(err, database.ErrFollowDoesNotExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			ctx.Logger.WithError(err).WithField("id", id).Error("can't retrieve the follow")
		}
		return
	}

	follow.FollowFromDatabase(datafollow)

	err = rt.db.RemoveFollow(follow.FollowToDatabase())
	if errors.Is(err, database.ErrFollowDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("id", id).Error("can't remove the follow")
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
