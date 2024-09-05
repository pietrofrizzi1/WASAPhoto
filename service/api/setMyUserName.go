package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	username := ps.ByName("singleusername")

	// Decodifica il corpo della richiesta
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))
	user.Id = token

	// Imposta il nome utente nel database
	dbuser, err := rt.db.SetUsername(user.ConvertForDatabase(), username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Rispondi con JSON usando respondWithJSON
	respondWithJSON(w, http.StatusCreated, user)
}
