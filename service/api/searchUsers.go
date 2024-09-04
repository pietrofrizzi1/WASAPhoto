package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	fmt.Println("Ciao, mondo!")
	var requestUser User
	var userList []User
	token := getAuthorization(r.Header.Get("Authorization"))
	requestUser.Id = token

	// Verifica dell'utente richiedente tramite il token
	dbrequestuser, err := rt.db.CheckUserById(requestUser.CovertForDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requestUser.ConvertForApplication(dbrequestuser)

	// Recupera il parametro username dalla URL
	username := ps.ByName("searchedusername")

	// Cerca utenti con nomi simili nel database
	dbusers, err := rt.db.SearchUsersByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Converti i risultati del database in un formato adatto per il JSON
	for _, dbuser := range dbusers {
		var user User
		user.ConvertForApplication(dbuser)
		userList = append(userList, user)
	}

	// Imposta l'intestazione del contenuto come JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica l'elenco degli utenti come JSON e scrivilo nella risposta
	_ = json.NewEncoder(w).Encode(userList)
}
