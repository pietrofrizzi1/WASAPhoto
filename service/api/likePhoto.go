package api


type Like struct {
	Username string
	
}

var Likes = [][]Like{
  []Like{
	
	Like{
		Username: "",
		
	},
	Like{
		Username: "",
		
	},
	Like{
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
nuovoLike := r.URL.Query().Get("immagine")
nuovoLike:= Like{
	Username:nuovoLike,
}
Likes= append(Likes,nuovoLike)
json.NewEncoder(w).Encode(nuovoLike)



}
