package api

type Followed struct {
	UserFollowing string
	Usernameofthefollowed string
	
	
}

type Like struct {
	Username string
	
}
type Comment struct {
	Username string
	Text string
	
}

type Photo struct {
	Username string
	Image string
    Likes
	Comments
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
nomeUser := r.URL.Query().Get("user da mostrare")
var fotoutente = []Photo{}
var seguitiutente=[]Followed{}

for _, photo := range Photo {
	if strings.EqualFold(nomeUser, photo.Username) {
		fotoutente=append(fotoutente,photo)
	
		
	}
}

for _, follow := range Followed {
	if strings.EqualFold(nomeUser ,follow.UserFollowing) {
		seguitiutente=append(FollowedUsers,follow)
	
		
	}
}
json.NewEncoder(w).Encode(Photos)
json.NewEncoder(w).Encode(Followed)



}
