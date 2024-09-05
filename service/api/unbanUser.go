package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var ban Ban
	var user User

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))

	// Converti l'ID dell'utente bannato
	id, err := strconv.ParseUint(ps.ByName("singlebanneduser"), 10, 64)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	ban.BanId = id

	// Recupera l'username dalla URL e ottieni i dettagli dell'utente
	username := ps.ByName("singleusername")
	user.Username = username
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	ban.UserId = token
	ban.BannedId = user.Id

	// Rimuovi il ban dal database
	if err := rt.db.RemoveBan(ban.BanConvertForDatabase()); err != nil {
		if errors.Is(err, database.ErrBanNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound) // Gestione dell'errore specifico
		} else {
			ctx.Logger.WithError(err).WithField("id", id).Error("can't remove the ban")
			handleError(w, err) // Gestione degli altri errori
		}
		return
	}

	// Rispondi con status No Content
	w.WriteHeader(http.StatusNoContent)
}
