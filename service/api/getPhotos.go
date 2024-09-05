package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var requestUser User
	var photoList database.Photos

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))

	// Imposta l'ID dell'utente che fa la richiesta
	requestUser.Id = token

	// Controlla se l'utente che fa la richiesta esiste nel database
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	requestUser.ConvertForApplication(dbrequestuser)

	// Ottieni l'ID dell'utente di cui vogliamo le foto
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Ottieni le foto dell'utente
	photos, err := rt.db.GetPhotos(user.ConvertForDatabase(), token)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Imposta l'ID dell'utente che fa la richiesta e le foto
	photoList.RequestUser = requestUser.Id
	photoList.Identifier = token
	photoList.Photos = photos

	// Rispondi con JSON usando respondWithJSON
	w.Header().Set("Content-Type", "application/json") // Modificato il Content-Type per JSON
	respondWithJSON(w, http.StatusOK, photoList)
}
