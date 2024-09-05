package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

// PROVARE QUESTAAAAAAAAAAA
func (router *_router) banUser(w http.ResponseWriter, r *http.Request, params httprouter.Params, context reqcontext.RequestContext) {
	// Estrarre e validare il token
	authHeader := r.Header.Get("Authorization")
	token := getAuthorization(authHeader)

	// Ottenere l'username dal percorso
	username := params.ByName("singleusername")

	// Recuperare i dati dell'utente
	dbUser, err := router.db.GetUserId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user User
	user.ConvertForApplication(dbUser)

	// Estrarre e validare l'ID dell'utente bannato
	bannedUserIDStr := params.ByName("singlebanneduser")
	bannedUserID, err := strconv.ParseUint(bannedUserIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Creare un nuovo record di ban
	var ban Ban
	ban.BanId = bannedUserID
	ban.BannedId = user.Id
	ban.UserId = token

	dbBan, err := router.db.AddBan(ban.BanConvertForDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ban.BanConvertForApplication(dbBan)

	// Rimuovere commenti e like associati all'utente bannato
	if err := removeUserInteractions(router, token, user.Id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Rimuovere eventuali follow
	if _, err := router.db.GetFollowingId(token, user.Id); err == nil {
		if err := router.db.RemoveFollow(token, user.Id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Rispondere con il ban creato
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ban); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// removeUserInteractions rimuove commenti e like per un utente specifico
func removeUserInteractions(router *_router, token uint64, userID uint64) error {
	if err := router.db.RemoveComments(token, userID); err != nil {
		return err
	}
	if err := router.db.RemoveLikes(token, userID); err != nil {
		return err
	}
	return nil
}
