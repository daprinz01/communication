package main

import (
	"communication/controllers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(controllers.TrackResponseTime)
	// Add middleware to run before request
	api.Use(controllers.AuthorizationMiddleware)
	// Add handlers
	api.HandleFunc("/send/email", controllers.SendEmail).Methods(http.MethodPost)
	api.HandleFunc("/send/newsletter", controllers.SendNewsletter).Methods(http.MethodPost)
	api.HandleFunc("/send/sms", controllers.SendSMS).Methods(http.MethodPost)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8083",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// Configure Logging
	logFileLocation := os.Getenv("LOG_FILE_LOCATION")
	if logFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
		log.Println("Successfully initialized log file...")
	}
	go func() {
		log.Println("Starting Server...")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down...")
	os.Exit(0)
}
