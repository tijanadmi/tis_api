package main

import (
	"os"
	"testing"

	"github.com/tijanadmi/tis-api/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {

	app.DB = &dbrepo.TestDBRepo{}
	//app.models = models.NewModels(db)
	app.Domain = "example.com"
	app.JWTSecret = ""
	os.Exit(m.Run())
}
