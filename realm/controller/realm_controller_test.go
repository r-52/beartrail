package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/r-52/beartrail/realm/db"
)

func TestGetHostname(t *testing.T) {
	// dirty way
	db.Connect()

	req, err := http.NewRequest("GET", "/realm", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("hostname", "localhost")
	req.URL.RawQuery = q.Encode()

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRealm)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
