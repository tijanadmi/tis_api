package main

import (
	"os"
	"testing"
	"time"

	"github.com/tijanadmi/tis-api/repository/dbrepo"
)

var app application
var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleGFtcGxlLmNvbSIsImV4cCI6MTY4NDA0Njk1OCwiaWF0IjoxNjg0NDA2OTU4LCJpc3MiOiJleGFtcGxlLmNvbSIsIm5hbWUiOiJ0aWphbmEiLCJyb2xlcyI6W3siSUQiOjEsIklkVXNlciI6MCwiSWRSb2xlIjowLCJSb2xlQ29kZSI6IiIsIlJvbGVOYW1lIjoiTkRDIn0seyJJRCI6MiwiSWRVc2VyIjowLCJJZFJvbGUiOjAsIlJvbGVDb2RlIjoiIiwiUm9sZU5hbWUiOiJEV0gifV0sInN1YiI6IjEiLCJ0eXAiOiJKV1QifQ.M8OKU-_TJ-jbNuzdcmQJSf7_izkfvc45smiSbo06WDc"

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
