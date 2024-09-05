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
	token := getAuthorization(r.Header.Get("Authorization"))
	user.Username = ps.ByName("singleusername")
	dbuser, err := rt.db.CheckUserByUsername(user.ConvertForDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ConvertForApplication(dbuser)
	dban, err := rt.db.GetBan(user.ConvertForDatabase(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	ban.BanConvertForApplication(dban)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ban)
}
