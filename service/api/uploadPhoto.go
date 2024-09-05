package api

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var photo Photo

	// Estrai il token e l'username
	token := extractToken(r.Header.Get("Authorization"))
	user.Id = token
	user.Username = ps.ByName("singleusername")

	// Verifica l'utente
	dbuser, err := rt.db.CheckUser(user.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	user.ConvertForApplication(dbuser)

	// Recupera l'ID della foto
	id, err := strconv.ParseUint(ps.ByName("singlephoto"), 10, 64)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	photo.Id = id

	// Leggi il file dalla richiesta
	photo.File, err = io.ReadAll(r.Body)
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}

	// Imposta la data e l'ID dell'utente
	photo.Date = time.Now().Format("2006-01-02 15:04:05")
	photo.UserId = user.Id

	// Salva la foto nel database
	dbphoto, err := rt.db.SetPhoto(photo.PhotoConvertForDatabase())
	if err != nil {
		handleError(w, err) // Gestione dell'errore
		return
	}
	photo.PhotoConvertForApplication(dbphoto)

	// Rispondi con la foto appena creata
	respondWithJSON(w, http.StatusCreated, photo)
}
