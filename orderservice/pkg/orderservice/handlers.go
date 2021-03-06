package orderservice

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type MenuItem struct {
	id       string `json:"id"`
	quantity int    `json:"quantity"`
}

type Order struct {
	id        string     `json:"id"`
	menuitems []MenuItem `json:"menuitems"`
}
type GetOrder struct {
	id                 string     `json:"id"`
	menuitems          []MenuItem `json:"menuitems"`
	orderedAtTimestamp int        `json:"timestamp"`
	cost               int        `json:"cost"`
}

var item1 = MenuItem{
	id:       "3fa85f64-5717-4562-b3fc-2c963f66afa6",
	quantity: 0,
}
var order1 = []Order{
	{
		id:        "3fa85f642-57172-45622-b3fc2-2c963f66afa62",
		menuitems: []MenuItem{item1},
	},
}

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/hello-World", helloWorld).Methods(http.MethodGet)
	s.HandleFunc("/orders", getOrders).Methods(http.MethodGet)
	s.HandleFunc("/order/{:id}", getOrder).Methods(http.MethodGet)

	return logMiddleWare(r)
}

func logMiddleWare(h http.Handler) http.Handler {
	was := time.Now()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
		passedTime := time.Now().Sub(was)
		log.WithFields(log.Fields{
			"timePassed": passedTime,
		}).Info("time passed")
	})
}

func getOrders(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(order1)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write respons error")
	}
}
func getOrder(w http.ResponseWriter, r *http.Request) {

	var item = MenuItem{
		id:       getId(r),
		quantity: 1,
	}

	b, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(b))

}
func getParam(needingKey string, r *http.Request) string {
	return r.URL.Query().Get(needingKey)
}
func getId(r *http.Request) string {
	vars := mux.Vars(r)
	id := vars["id"]
	return id
}
func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello-world!")
}
