package main

import (
	"log"
	"os"
	"testing"

	"github.com/tijanadmi/tis-api/models"
)

var app application

func TestMain(m *testing.M) {
	var cfg config
	cfg.db.dsn = ""

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//app.DB = &dbrepo.TestDBRepo{}
	app.config = cfg
	app.models = models.NewModels(db)
	app.Domain = "example.com"
	app.JWTSecret = "verysecret"
	os.Exit(m.Run())
}
