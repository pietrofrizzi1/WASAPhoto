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
	var user User
	var like Like

	// Ottieni l'username dalla URL e recupera i dettagli dell'utente
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Converti gli ID della foto e del like
	photoid, err := strconv.ParseUint(ps.ByName("singlephoto"), 10, 64)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	likeid, err := strconv.ParseUint(ps.ByName("singlelike"), 10, 64)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	like.LikeId = likeid

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))
	like.UserIdentifier = token
	like.PhotoIdentifier = photoid
	like.PhotoOwner = user.Id

	// Rimuovi il like dal database
	if err := rt.db.RemoveLike(like.LikeConvertForDatabase()); err != nil {
		if errors.Is(err, database.ErrLikeNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound) // Gestione dell'errore specifico
		} else {
			ctx.Logger.WithError(err).WithField("id", likeid).Error("can't delete the like")
			handleError(w, err) // Gestione degli altri errori
		}
		return
	}

	// Rispondi con status No Content
	w.WriteHeader(http.StatusNoContent)
}
