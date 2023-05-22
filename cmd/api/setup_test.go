package main

import (
	"os"
	"testing"
	"time"

	"github.com/tijanadmi/tis-api/repository/dbrepo"
)

var app application
var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleGFtcGxlLmNvbSIsImV4cCI6MTY4NDM4MjUyNSwiaWF0IjoxNjg0NzQyNTI1LCJpc3MiOiJleGFtcGxlLmNvbSIsIm5hbWUiOiJ0aWphbmEiLCJyb2xlcyI6WyJOREMiLCJEV0giXSwic3ViIjoiMSIsInR5cCI6IkpXVCJ9.SGJTuUKeYfoG3fxJ42433GZ-qRzcgqSzmuTdLnAYRcY"

func TestMain(m *testing.M) {

	app.DB = &dbrepo.TestDBRepo{}
	//app.models = models.NewModels(db)
	app.Domain = "example.com"
	app.JWTSecret = ""
	app.JWTIssuer = "example.com"
	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}
	os.Exit(m.Run())
}
