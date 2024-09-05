package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like
	var user User

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))

	// Ottieni l'ID dell'utente dal database
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Ottieni e imposta l'ID della foto e dell'like
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
	like.UserIdentifier = token
	like.PhotoIdentifier = photoid
	like.PhotoOwner = user.Id

	// Salva il like nel database
	dblike, err := rt.db.SetLike(like.LikeConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	like.LikeConvertForApplication(dblike)

	// Rispondi con JSON usando respondWithJSON
	respondWithJSON(w, http.StatusCreated, like)
}
