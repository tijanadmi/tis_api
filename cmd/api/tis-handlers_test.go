package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_authenticate(t *testing.T) {
	var theTests = []struct {
		name               string
		requestBody        string
		expectedStatusCode int
	}{
		{"valid user", `{"username":"test","password":"test"}`, http.StatusAccepted},
		{"not json", `I'm not JSON`, http.StatusUnauthorized},
		{"empty json", `{}`, http.StatusUnauthorized},
		{"empty user", `{"username":""}`, http.StatusUnauthorized},
		{"empty password", `{"username":"test"}`, http.StatusUnauthorized},
		{"invalid user", `{"username":"test","password":"secret"}`, http.StatusUnauthorized},
	}

	for _, e := range theTests {

		req, _ := http.NewRequest("POST", "/auth", strings.NewReader(e.requestBody))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.authenticate)

		handler.ServeHTTP(rr, req)

		if e.expectedStatusCode != rr.Code {
			t.Errorf("%s: returned wrong status code; expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}
