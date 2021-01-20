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

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Create Server and Route Handlers

	// r := mux.NewRouter()
	// api := r.PathPrefix("/api/v1").Subrouter()
	// api.Use(controllers.TrackResponseTime)
	// // Add middleware to run before request
	// api.Use(controllers.AuthorizationMiddleware)
	// // Add handlers
	// api.HandleFunc("/send/email", controllers.SendEmail).Methods(http.MethodPost)
	// api.HandleFunc("/send/newsletter", controllers.SendNewsletter).Methods(http.MethodPost)
	// api.HandleFunc("/send/sms", controllers.SendSMS).Methods(http.MethodPost)
	srv := &http.Server{

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
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: &lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		},
	}))
	e.Use(controllers.TrackResponseTime)
	// e.Use(middleware.CSRF())
	e.Use(middleware.Recover())
	// Enable metrics middleware
	e.Use(echoPrometheus.MetricsMiddleware())
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	// e.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	api := e.Group("/api/v1")
	api.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == os.Getenv("CLIENT_ID"), nil
	}))
	api.POST("/send/email", controllers.SendEmail)
	api.POST("/send/newsletter", controllers.SendNewsletter)
	api.POST("/send/sms", controllers.SendSMS)

	// go func() {
	// 	log.Println("Starting Server...")
	// 	if err := srv.ListenAndServe(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	go func() {
		log.Println("Starting Server...")
		e.Logger.Fatal(e.StartServer(srv))
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
