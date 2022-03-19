package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"persianblack.com/communication/controllers"

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

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
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: &lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default,
		},
		Format: "{\"@timestamp\":\"${time_rfc3339}\", \"uri\":\"${uri}\", \"remote_ip\":\"${remote_ip}\", \"host\":\"${host}\", \"id\":\"${id}\", \"method\":\"${method}\", \"user_agent\":\"${user_agent}\", \"status\":\"${status}\", \"error\":\"${error}\", \"latency\":\"${latency}\", \"latency_human\":\"${latency_human}\", \"bytes_in\":\"${bytes_in}\", \"bytes_out\":\"${bytes_out}\", \"message\":\"Echo http request logger\", \"microservice\": \"persian.black.communication.service\", \"level\":\"info\", \"user_agent\":\"${user_agent}\"}",
	}))
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	// e.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	api := e.Group("/api/v1")
	api.Use(controllers.Authorize)
	api.POST("/send/email", controllers.SendEmail)
	api.POST("/send/newsletter", controllers.SendNewsletter)
	api.POST("/send/sms", controllers.SendSMS)

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
