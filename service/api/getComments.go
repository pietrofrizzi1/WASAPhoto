package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var requestUser User
	var photo Photo
	var commentList database.Comments

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))
	requestUser.Id = token

	// Controlla se l'utente esiste nel database
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	requestUser.ConvertForApplication(dbrequestuser)

	// Estrai l'ID della foto
	photoid, err := parseUserID(ps.ByName("singlephoto")) // Usa parseUserID per l'ID della foto
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

	// Recupera l'utente tramite username
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Recupera i commenti della foto
	comments, err := rt.db.GetComments(photo.Id)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Imposta i dettagli della lista dei commenti
	commentList.RequestIdentifier = requestUser.Id
	commentList.PhotoIdentifier = photo.Id
	commentList.PhotoOwner = user.Id
	commentList.Comments = comments

	// Rispondi con JSON usando respondWithJSON
	respondWithJSON(w, http.StatusOK, commentList)
}
