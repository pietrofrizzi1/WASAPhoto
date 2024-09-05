package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var userList []User

	// Estrai il token usando extractToken
	token := extractToken(r.Header.Get("Authorization"))

	// Verifica dell'utente richiedente tramite il token
	var requestUser User
	requestUser.Id = token
	dbrequestuser, err := rt.db.CheckUserById(requestUser.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	requestUser.ConvertForApplication(dbrequestuser)

	// Recupera il parametro username dalla URL
	username := ps.ByName("searchedusername")

	// Cerca utenti con nomi simili nel database
	dbusers, err := rt.db.SearchUsersByUsername(username)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Converti i risultati del database in un formato adatto per il JSON
	for _, dbuser := range dbusers {
		var user User
		user.ConvertForApplication(dbuser)
		userList = append(userList, user)
	}

	// Rispondi con JSON usando respondWithJSON
	respondWithJSON(w, http.StatusOK, userList)
}
