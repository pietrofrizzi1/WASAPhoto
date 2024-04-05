package api

type Followed struct {
	Username string
	
	
}

var FollowedUsers = [][]Followed{
  []Followed{
	
	Followed{
		Username: "",
		
		
	},
	Followed{
		Username: "",
		
		
	},
	Followed{
		Username: "",
		
		
	},
  },

}

import (
"github.com/julienschmidt/httprouter"
"net/http"
"encoding/json"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
w.Header().Set("content-type", "application/json")
nuovoSeguito := r.URL.Query().Get("seguito")
nuovoSeguito:=Followed{
	Username:nuovoSeguito,
	
}
FollowedUsers= append(FollowedUsers,nuovoSeguito)
json.NewEncoder(w).Encode(nuovoSeguito)



}

