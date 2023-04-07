package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/r-52/beartrail/realm/controller"
	"github.com/r-52/beartrail/realm/db"
	"github.com/r-52/beartrail/realm/models"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/health"))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Get("/realm", controller.GetRealm)

	prepareDbAndMigrate()
	http.ListenAndServe(":3100", r)
}

func prepareDbAndMigrate() {
	db.Connect()
	db.Database.AutoMigrate(&models.Realm{})

	localhost, err := models.GetRealmByHostname("localhost")
	if err != nil {
		localhost = models.Realm{
			Hostname: "localhost",
		}
		db.Database.Create(&localhost)
	}

}
