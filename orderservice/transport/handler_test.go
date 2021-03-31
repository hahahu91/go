package orderservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrders(t *testing.T) {
	const RequestAddress = "http://localhost:8000/api/v1/orders"
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", RequestAddress, nil)

	getOrders(w, r)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code iw wrong. Have: %d. want:%d", response.StatusCode, http.StatusOK)
	}
	jsonSting, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items := make([]MenuItem, 10)
	if err = json.Unmarshal(jsonSting, &items); err != nil {
		t.Errorf("Can't parse json response with errpr %v", err)
	}
	fmt.Println("Test passed with response: " + string(jsonSting))
}

