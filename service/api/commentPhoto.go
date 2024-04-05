package api

type Comment struct {
	Username string
	Text string
	
}

var Comments = [][]Comment{
  []Comment{
	
	Comment{
		Username: "",
		Text string,
		
	},
	Comment{
		Username: "",
		Text string,
		
	},
	Comment{
		Username: "",
		Text string,
		
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
nuovoCommento := r.URL.Query().Get("commento")
nuovoCommento:= Comment{
	Username:nuovoLike,
	Text:"",
}
Comments= append(Comments,nuovoCommento)
json.NewEncoder(w).Encode(nuovoCommento)



}

