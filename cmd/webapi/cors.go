package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func applyCORSHandler(h http.Handler) http.Handler {
	log.Print("1")
	http.HandleFunc("/test", testHandler)

	log.Print("Server listening on port 3000")
	http.ListenAndServe(":3000", nil)
	return handlers.CORS(

		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-Requested-With"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"}),
		handlers.AllowedOrigins([]string{"*"}), // Permetti tutte le origini, personalizza se necessario
		handlers.MaxAge(1),
	)(h)
}
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is up and running"))
}
