package api

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
var SeguitidaUser=[]Followed{}
for _, followed := range FollowedUsers{
	if strings.EqualFold(followed.UserFollowing,nomeUser) {
		SeguitidaUser=append(SeguitidaUser,followed)
		
for _, photo := range Photo {
	for _, seguito := range SeguitidaUser{
		if strings.EqualFold(seguito.Usernameofthefollowed, photo.Username) {
			fotoutente=append(fotoutente,photo)
		
			
		}
	}
	
}


json.NewEncoder(w).Encode(fotoutente)




}