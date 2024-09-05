package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var (
		follow   Follow
		user     User
		username = ps.ByName("singleusername")
	)

	// Recupera l'ID dell'utente
	dbUser, err := rt.db.GetUserId(username)
	if err != nil {
		handleError(w, err)
		return
	}
	user.ConvertForApplication(dbUser)

	// Ottieni l'ID dell'utente da seguire
	followedUserID, err := parseUserID(ps.ByName("singlefolloweduser"))
	if err != nil {
		handleError(w, err)
		return
	}

	// Prepara i dati per il follow
	follow.FollowId = followedUserID
	follow.FollowedId = user.Id
	follow.UserId = extractToken(r.Header.Get("Authorization"))

	// Salva la relazione di follow nel database
	dbFollow, err := rt.db.SetFollow(follow.FollowConvertForDatabase())
	if err != nil {
		handleError(w, err)
		return
	}
	follow.FollowConvertForApplication(dbFollow)

	// Invia la risposta
	respondWithJSON(w, http.StatusCreated, follow)
}
