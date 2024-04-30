package main

import (
	"context"
	"joi-energy-golang/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})
}

func main() {
	Run()
}

// Run starts the HTTP server
func Run() {
	server := router.NewServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		println()
		log.Info("Shutting down server...")

		err := gracefulShutdown(server, 25*time.Second)

		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Fatal("Server stopped!")
		}

		os.Exit(0)
	}()

	log.WithFields(log.Fields{
		"port": server.Addr,
	}).Info("Starting server...")

	log.Fatal(server.ListenAndServe())
}

func gracefulShutdown(server *http.Server, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return server.Shutdown(ctx)
}
