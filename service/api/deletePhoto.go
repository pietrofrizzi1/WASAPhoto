package api

type Photo struct {
	Image string
}

var Photos = []Photo{
	
	Photo{
		Image: "Andrea3",
	},
	Photo{
		Image: "Luca!",
	},
	Photo{
		Image: "Edoardo1102",
	},

}

import (
"github.com/julienschmidt/httprouter"
"net/http"
"encoding/json"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
w.Header().Set("content-type", "application/json")
nuovaImmagine := r.URL.Query().Get("immagine")
nuovaImmagine:= Photo{
	Image:nuovaImmagine,
}
found := false
	for id, immagine := range Photos {
		if strings.EqualFold(nuovaImmagine, immagine.Image) {
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
Photos= append(Photos,nuovaImmagine)
json.NewEncoder(w).Encode(nuovaImmagine)

}