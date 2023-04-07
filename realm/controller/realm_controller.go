package controller

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/r-52/beartrail/realm/models"
)

func GetRealm(w http.ResponseWriter, r *http.Request) {
	hostname := r.URL.Query().Get("hostname")

	if len(hostname) == 0 {
		w.WriteHeader(404)
		return
	}

	realm, err := models.GetRealmByHostname(hostname)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	// respond to the client
	render.JSON(w, r, realm)
}
