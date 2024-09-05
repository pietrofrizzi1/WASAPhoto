package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var follow Follow

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))

	// Imposta lo username dell'utente
	user.Username = ps.ByName("singleusername")

	// Controlla se l'utente esiste nel database
	dbuser, err := rt.db.CheckUserByUsername(user.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Recupera i follower dell'utente
	dbfollow, err := rt.db.GetFollowers(user.ConvertForDatabase(), token)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Converte i dati dei follower per l'applicazione
	follow.FollowConvertForApplication(dbfollow)

	// Rispondi con JSON usando respondWithJSON
	respondWithJSON(w, http.StatusOK, follow)
}
