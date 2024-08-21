package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	id, err := strconv.ParseUint(ps.ByName("commentid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datacomment, err := rt.db.GetCommentById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.CommentFromDatabase(datacomment)
	err = rt.db.RemoveComment(comment.CommentToDatabase())
	if errors.Is(err, database.ErrCommentDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
