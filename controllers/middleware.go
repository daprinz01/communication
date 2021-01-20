package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// RecoverWrap helps to recover from a panic. Currently not in use simply because it doesn't work!
// func RecoverWrap(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var err error
// 		defer func() {
// 			r := recover()
// 			if r != nil {
// 				switch t := r.(type) {
// 				case string:
// 					err = errors.New(t)
// 				case error:
// 					err = t
// 				default:
// 					err = errors.New("Unknown error")
// 				}
// 				sendMeMail(err)
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 			}
// 		}()
// 		h.ServeHTTP(w, r)
// 	})
// }

// func sendMeMail(err error) {
// 	// send mail
// 	log.Println("Error in sendMeMail: ")
// 	log.Println(err)
// }

// AuthorizationMiddleware is used to authorize API calls
// func AuthorizationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		w.Header().Set("Content-Type", "application/json")
// 		client := r.Header.Get("clientId")
// 		approvedClientID := os.Getenv("CLIENT_ID")
// 		// var err error
// 		var errorResponse models.ErrorResponse

// 		if client != approvedClientID {
// 			log.Println(fmt.Sprintf("Unauthorised request from client: %s", approvedClientID))
// 			errorResponse.Errorcode = "01"
// 			errorResponse.ErrorMessage = "Unauthorised"

// 			response, err := json.MarshalIndent(errorResponse, "", "")
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write(response)
// 			return
// 		}
// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 		next.ServeHTTP(w, r)
// 	})
// }

// TrackResponseTime is used to track the response time of api calls
func TrackResponseTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Measure response time
		start := time.Now()
		if err := next(c); err != nil {
			c.Error(err)
		}
		responseTime := time.Since(start)

		// Write it to the log
		log.Println(fmt.Sprintf("Request executed in %v", responseTime))
		return nil
	}

}
