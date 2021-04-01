package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"orderservice/pkg/orderservice"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}
	serverUrl := ":8000"
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")

	killSignalChan := getKillSignalChan()
	srv := startServer(serverUrl)

	waitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())
}
func startServer(serverUrl string) *http.Server {
	router := orderservice.Router()
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	return srv
}
func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(KillSignalChan <-chan os.Signal) {
	KillSignal := <-KillSignalChan
	switch KillSignal {
	case os.Interrupt:
		log.Info("got SIGINT,,,")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
