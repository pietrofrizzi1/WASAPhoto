package api

type Game struct {
	Id int
	secret int
	Outcome string
	Guesses int
}

var Games = []Game{
	Game{
		Id: 0,
		secret: 7,
		Outcome: "win",
		Guesses: 5,
	},
}

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
	UserAccessing := r.URL.Query().Get("username")
	found := false
	for _, user := range Users {
		if strings.EqualFold(UserAccessing, user.Username) {
			found = true
			break}
		}
	if found{
		json.NewEncoder(w).Encode(UserAccessing)
	}  
	else{
		NuovoUser:= User{
			Username:UserAccessing,
		}
		Users= append(Users,NuovoUser)
        json.NewEncoder(w).Encode(UserAccessing)
	}

}