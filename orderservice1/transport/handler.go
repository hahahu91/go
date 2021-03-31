package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/hello-world", helloWorld).Methods(http.MethodGet)
	s.HandleFunc("/cat", getOrder).Methods(http.MethodGet)

	s.HandleFunc("/order", createOrder).Methods(http.MethodPost)
	return logMiddleware(r)
}

type MenuItem struct {
	id       string `json:"id"`
	quantity int    `json: "quantity"`
}

type Order struct {
	id        string     `json:"id`
	menuItems []MenuItem `json:"MenuItems"`
}

func getOrder(w http.ResponseWriter, _ *http.Request) {
	var item = MenuItem{
		id:       "17321-31231ads-dai123",
		quantity: 2,
	}
	var order = Order{
		id:        "324",
		menuItems: []MenuItem{item},
	}
	//order.id = "324"
	//order.menuItems = item1

	//cat.Name = "Kot"
	b, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("Write response error")
	}

}

//
func createOrder(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
	var msg Order
	err = json.Unmarshal(b, &msg)
	if err != nil {

	}
}
func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
func logMiddleware(h http.Handler) http.Handler {
	was := time.Now()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
		now := time.Now()
		passedTime := now.Sub(was)

		log.WithFields(log.Fields{
			"timePassed": passedTime,
		}).Info("time passed")
	})
}
