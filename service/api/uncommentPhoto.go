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
	var user User
	var comment Comment

	// Converti l'ID della foto e del commento
	photoid, err := strconv.ParseUint(ps.ByName("singlephoto"), 10, 64)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	commentid, err := strconv.ParseUint(ps.ByName("singlecomment"), 10, 64)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	comment.Id = commentid
	comment.PhotoId = photoid

	// Ottieni l'username dalla URL e recupera i dettagli dell'utente
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))
	comment.UserId = token
	comment.PhotoOwner = user.Id

	// Rimuovi il commento dal database
	if err := rt.db.RemoveComment(comment.CommentConvertForDatabase()); err != nil {
		if errors.Is(err, database.ErrCommentNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound) // Gestione dell'errore specifico
		} else {
			ctx.Logger.WithError(err).WithField("id", commentid).Error("can't delete the comment")
			handleError(w, err) // Gestione degli altri errori
		}
		return
	}

	// Rispondi con status No Content
	w.WriteHeader(http.StatusNoContent)
}
