package api

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	dbFollow, err := rt.db.SetFollow(follow.FollowCovertForDatabase())
	if err != nil {
		handleError(w, err)
		return
	}
	follow.FollowConvertForApplication(dbFollow)

	// Invia la risposta
	respondWithJSON(w, http.StatusCreated, follow)
}

// handleError gestisce gli errori di risposta HTTP
func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// parseUserID converte una stringa in un ID utente
func parseUserID(idStr string) (uint64, error) {
	return strconv.ParseUint(idStr, 10, 64)
}

// extractToken estrae il token dal header Authorization
func extractToken(authHeader string) uint64 {
	return getAuthorization(authHeader)
}

// respondWithJSON invia una risposta JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
