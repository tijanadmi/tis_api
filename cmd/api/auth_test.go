package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_app_getTokenFromHeaderAndVerify(t *testing.T) {
	roles := []string{"NDC", "DWH"}
	testUser := jwtUser{
		ID:        1,
		FirstName: "tijana",
		LastName:  "tijana",
		Roles:     roles,
	}

	tokens, _ := app.auth.GenerateTokenPair(&testUser)

	var tests = []struct {
		name          string
		token         string
		errorExpected bool
		setHeader     bool
		issuer        string
		role          string
	}{
		{"valid", fmt.Sprintf("Bearer %s", tokens.Token), false, true, app.Domain, "NDC"},
		{"valid expired", fmt.Sprintf("Bearer %s", expiredToken), true, true, app.Domain, "NDC"},
		{"no header", "", true, false, app.Domain, "NDC"},
		{"invalid token", fmt.Sprintf("Bearer %s1", tokens.Token), true, true, app.Domain, "NDC"},
		{"no bearer", fmt.Sprintf("Bear %s1", tokens.Token), true, true, app.Domain, "NDC"},
		{"three header parts", fmt.Sprintf("Bearer %s 1", tokens.Token), true, true, app.Domain, "NDC"},
		{"wrong role", fmt.Sprintf("Bearer %s", tokens.Token), true, true, app.Domain, "RDC"},
		// make sure the next test is the last one to run
		{"wrong issuer", fmt.Sprintf("Bearer %s", tokens.Token), true, true, "anotherdomain.com", "NDC"},
	}

	for _, e := range tests {
		if e.issuer != app.auth.Issuer {
			app.auth.Issuer = e.issuer
			tokens, _ = app.auth.GenerateTokenPair(&testUser)
		}
		req, _ := http.NewRequest("GET", "/", nil)
		if e.setHeader {
			req.Header.Set("Authorization", e.token)
		}

		rr := httptest.NewRecorder() //response recorder

		_, _, err := app.auth.GetTokenFromHeaderAndVerify(rr, req, e.role)
		if err != nil && !e.errorExpected {
			t.Errorf("%s: did not expect error, but got one - %s", e.name, err.Error())
		}

		if err == nil && e.errorExpected {
			t.Errorf("%s: expected error, but did not get one", e.name)
		}

		app.auth.Issuer = "example.com"
	}
}
