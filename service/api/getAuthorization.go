package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

func getAuthorization(message string) uint64 {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	stringToken := re.FindAllString(message, -1)
	token, _ := strconv.Atoi(stringToken[0])
	return uint64(token)
}
func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// parseUserID converte una stringa in un ID utente
func parseUserID(idStr string) (uint64, error) {
	return strconv.ParseUint(idStr, 10, 64)
}

// extractToken estrae il token dal header Authorization
func extractToken(authHeader string) uint64 {
	return getAuthorization(authHeader)
}

// respondWithJSON invia una risposta JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
