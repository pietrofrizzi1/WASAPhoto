package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
	"github.com/pietrofrizzi1/WASAPhoto/service/database"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var ban Ban
	id, err := strconv.ParseUint(ps.ByName("banid"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	databan, err := rt.db.GetBanById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ban.BanFromDatabase(databan)
	err = rt.db.RemoveBan(ban.BanToDatabase())
	if errors.Is(err, database.ErrBanDoesNotExist) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("The ban can't be removed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
