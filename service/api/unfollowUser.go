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
	var user User

	// Ottieni l'username dalla URL e recupera i dettagli dell'utente
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Converti l'ID dell'utente seguito
	id, err := strconv.ParseUint(ps.ByName("singlefolloweduser"), 10, 64)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	follow.FollowId = id
	follow.FollowedId = user.Id

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))
	follow.UserId = token

	// Rimuovi il follow dal database
	if err := rt.db.RemoveFollow(follow.UserId, follow.FollowedId); err != nil {
		if errors.Is(err, database.ErrBanNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound) // Gestione dell'errore specifico
		} else {
			ctx.Logger.WithError(err).WithField("id", id).Error("can't delete the follow")
			handleError(w, err) // Gestione degli altri errori
		}
		return
	}

	// Rispondi con status No Content
	w.WriteHeader(http.StatusNoContent)
}
