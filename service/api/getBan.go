package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pietrofrizzi1/WASAPhoto/service/api/reqcontext"
)

func (rt *_router) getBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User
	var ban Ban

	// Estrai il token usando la funzione extractToken
	token := extractToken(r.Header.Get("Authorization"))

	// Imposta lo username dall'URL
	user.Username = ps.ByName("singleusername")

	// Controlla se l'utente esiste nel database
	dbuser, err := rt.db.CheckUserByUsername(user.ConvertForDatabase())
	if err != nil {
		handleError(w, err) // Usa handleError per gestire gli errori
		return
	}

	// Converti l'utente per l'applicazione
	user.ConvertForApplication(dbuser)

	// Recupera il ban dal database
	dban, err := rt.db.GetBan(user.ConvertForDatabase(), token)
	if err != nil {
		handleError(w, err) // Usa handleError per gestire gli errori
		return
	}

	// Converti il ban per l'applicazione
	ban.BanConvertForApplication(dban)

	// Rispondi con JSON usando la funzione respondWithJSON
	respondWithJSON(w, http.StatusOK, ban)
}
