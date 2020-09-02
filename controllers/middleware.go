package controllers

import (
	"communication/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// AuthorizationMiddleware is used to authorize API calls
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		client := r.Header.Get("clientId")
		approvedClientID := os.Getenv("CLIENT_ID")
		// var err error
		var errorResponse models.ErrorResponse

		if client != approvedClientID {
			log.Println(fmt.Sprintf("Unauthorised request from client: %s", approvedClientID))
			errorResponse.Errorcode = "01"
			errorResponse.ErrorMessage = "Unauthorised"

			response, err := json.MarshalIndent(errorResponse, "", "")
			if err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
