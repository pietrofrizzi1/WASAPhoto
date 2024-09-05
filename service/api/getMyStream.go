package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photoList database.Stream

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))

	// Imposta l'ID e il nome utente dell'utente
	username := ps.ByName("singleusername")
	user.Id = token
	user.Username = username

	// Controlla se l'utente esiste nel database
	dbuser, err := rt.db.CheckUser(user.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Ottieni lo stream dell'utente
	photos, err := rt.db.GetMyStream(user.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Imposta l'ID dell'utente e le foto nello stream
	photoList.Identifier = token
	photoList.Photos = photos

	// Rispondi con JSON usando respondWithJSON
	respondWithJSON(w, http.StatusOK, photoList)
}
