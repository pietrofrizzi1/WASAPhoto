package api

type User struct {
	Username string
}

var Users = []User{
	
		User{
			Username: "Andrea3",
		},
		User{
			Username: "Luca!",
		},
		User{
            Username: "Edoardo1102",
		},
	
}

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	UserAccessing, NomeCambiato := r.URL.Query().Get("username")
	found := false
	for i, user := range Users {
		if strings.EqualFold(UserAccessing, user.Username) {
			found = i
			break}

		}
	if found{
		
		Users[found].Username = NomeCambiato
		json.NewEncoder(w).Encode(NomeCambiato)
	}  
	else{
        w.WriteHeader(http.StatusNotFound)
		return
	}

}