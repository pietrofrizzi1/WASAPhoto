package api

import (
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) getBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var ban Ban
	token := getToken(r.Header.Get("Authorization"))
	user.Username = ps.ByName("singleusername")
	dbuser, err := rt.db.CheckUserByUsername(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	dban, err := rt.db.GetBan(user.ToDatabase(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	ban.BanFromDatabase(dban)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ban)
}
