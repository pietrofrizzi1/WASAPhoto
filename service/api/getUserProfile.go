package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var requestUser User
	var profile Profile

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

	// Ottieni l'ID dell'utente di cui vogliamo il profilo
	username := ps.ByName("singleusername")
	dbuser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Imposta i dettagli del profilo
	profile.RequestId = token
	profile.Id = user.Id
	profile.Username = user.Username

	// Ottieni e imposta il numero di follower
	followersCount, err := rt.db.GetFollowersCount(user.Id)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	profile.FollowersCount = followersCount

	// Ottieni e imposta il numero di following
	followingCount, err := rt.db.GetFollowingsCount(user.Id)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	profile.FollowingCount = followingCount

	// Ottieni e imposta lo stato di ban
	profile.BanStatus, err = rt.db.GetBanStatus(requestUser.Id, user.Id)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Ottieni e imposta lo stato di follow
	profile.FollowStatus, err = rt.db.GetFollowStatus(requestUser.Id, user.Id)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Controlla se l'utente è stato bannato
	profile.CheckIfBanned, err = rt.db.CheckIfBanned(requestUser.Id, user.Id)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Rispondi con JSON usando respondWithJSON
	respondWithJSON(w, http.StatusOK, profile)
}
