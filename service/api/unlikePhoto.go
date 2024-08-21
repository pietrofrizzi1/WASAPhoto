package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like
	id, err := strconv.ParseUint(ps.ByName("likeid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datalike, err := rt.db.GetLikeById(id)
	if err != nil {
		if errors.Is(err, database.ErrLikeDoesNotExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			ctx.Logger.WithError(err).WithField("id", id).Error("can't retrieve the like")
		}
		return
	}

	like.LikeFromDatabase(datalike)

	err = rt.db.RemoveLike(like.LikeToDatabase())
	if errors.Is(err, database.ErrLikeDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("id", id).Error("can't remove the like")
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
