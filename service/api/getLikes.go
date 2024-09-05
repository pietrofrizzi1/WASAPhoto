package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photo Photo
	var requestUser User

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))

	// Imposta l'ID dell'utente richiesto
	requestUser.Id = token

	// Controlla se l'utente esiste nel database
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	requestUser.ConvertForApplication(dbrequestuser)

	// Ottieni l'ID dell'utente tramite il nome utente
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Recupera l'ID della foto dalla richiesta
	photoid, err := strconv.ParseUint(ps.ByName("singlephoto"), 10, 64)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	photo.Id = photoid

	// Controlla se la foto esiste nel database
	dbphoto, err := rt.db.CheckPhoto(photo.PhotoConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	photo.PhotoConvertForApplication(dbphoto)

	// Ottieni i "like" della foto
	like, err := rt.db.GetLike(photo.Id, token)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Rispondi con JSON usando respondWithJSON
	respondWithJSON(w, http.StatusOK, like)
}
