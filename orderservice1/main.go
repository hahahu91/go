package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"orderservice/transport"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}
	c, err := parseEnv()
	serverUrl := c.ServeRESTAddress

	log.WithFields(log.Fields{"url": serverUrl}).Info("Starting the server")

	killSignalChan := getKillSignalChan()
	srv := startServer(serverUrl)
	waitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())
}

type Order struct {
	id       string
	quantity int
}

func getOrder(w http.ResponseWriter, _ *http.Request) {
	cat := Order{"Kot", 1}
	b, _ := json.Marshal(cat)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(b))
}
func startServer(serverUrl string) *http.Server {
	router := transport.Router()
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

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM....")
	}
}
