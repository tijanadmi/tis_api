package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
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

func Test_refreshToken(t *testing.T) {
	var tests = []struct {
		name               string
		token              string
		expectedStatusCode int
		resetRefreshTime   bool
	}{
		{"valid", "", http.StatusOK, true},
		/*{"valid but not yet ready to expire", "", http.StatusTooEarly, false},*/
		/*{"expired token", expiredToken, http.StatusUnauthorized, false},*/
	}

	roles := []string{"NDC", "DWH"}
	testUser := jwtUser{
		ID:        1,
		FirstName: "tijana",
		LastName:  "tijana",
		Roles:     roles,
	}

	oldRefreshTime := app.auth.RefreshExpiry

	for _, e := range tests {
		var tkn string
		if e.token == "" {
			if e.resetRefreshTime {
				app.auth.RefreshExpiry = time.Second * 1
			}
			tokens, _ := app.auth.GenerateTokenPair(&testUser)
			tkn = tokens.RefreshToken
		} else {
			tkn = e.token
		}

		postedData := url.Values{
			"refresh_token": {tkn},
		}

		req, _ := http.NewRequest("POST", "/refresh-token", strings.NewReader(postedData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.refreshToken)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s: expected status of %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		app.auth.RefreshExpiry = oldRefreshTime
	}

}

func Test_PiPiDDNHandlers(t *testing.T) {
	var tests = []struct {
		name           string
		method         string
		json           string
		paramID        string
		handler        http.HandlerFunc
		expectedStatus int
	}{
		{"getAllPiPiDDN", "GET", "", "", app.getAllPiPiDDN, http.StatusOK},
		{"deletePiPiDDN", "DELETE", "", "1", app.deletePiPiDDN, http.StatusAccepted},
		{"getPiPiDDNByID valid", "GET", "", "1", app.getPiPiDDNByID, http.StatusOK},
		{"getPiPiDDNByID invalid", "GET", "", "100", app.getPiPiDDNByID, http.StatusBadRequest},
		{"getPiPiDDNByID bad URL param", "GET", "", "Y", app.getPiPiDDNByID, http.StatusBadRequest},
		{
			"updatePiPiDDNIsklj valid",
			"PATCH",
			`{"dat_smene": "07.04.2023",
			"id_s_mrc": "8",
			 "tip_man": "1",
			 "id_tipob": "4",
			 "ob_id": "105",
			 "trafo_id": "",
			 "vrepoc": "07.04.2023 01:35:00",
			 "vrezav": "07.04.2023 02:25:00",
			 "id_s_grraz": "3",
			 "id_s_razlog": "21",
			 "man_tekst": "Od 23.03.23 07:45 na snazi planirani radovi broj: R-67/1 (AKZ metalnih nosača VN opreme u TRP 400kV od TR1 u  TS Beograd 8 Vasiljević Aleksandar). NA SNAZI;",
			 "id_s_nap": "",
			 "p2_traf_id": "",
			 "kor_uneo": "pi1100",
			 "ed_id": "1"}`,
			"",
			app.updatePiPiDDNIsklj,
			http.StatusAccepted,
		},
		{
			"updatePiPiDDNIsklj invalid",
			"PATCH",
			`{"dat_smene": "07.04.2023",
			"id_s_mrc": "8",
			 "tip_man": "1",
			 "id_tipob": "4",
			 "ob_id": "105",
			 "trafo_id": "",
			 "vrepoc": "07.04.2023 01:35:00",
			 "vrezav": "07.04.2023 02:25:00",
			 "id_s_grraz": "3",
			 "id_s_razlog": "21",
			 "man_tekst": "Od 23.03.23 07:45 na snazi planirani radovi broj: R-67/1 (AKZ metalnih nosača VN opreme u TRP 400kV od TR1 u  TS Beograd 8 Vasiljević Aleksandar). NA SNAZI;",
			 "id_s_nap": "",
			 "p2_traf_id": "",
			 "kor_uneo": "pi1100",
			 "ed_id": "100"}`,
			"",
			app.updatePiPiDDNIsklj,
			http.StatusBadRequest,
		},
		{
			"updatePiPiDDNIsklj invalid json",
			"PATCH",
			`{"dat_smene": "07.04.2023",
			id_s_mrc: "8",
			 "tip_man": "1",
			 "id_tipob": "4",
			 "ob_id": "105",
			 "trafo_id": "",
			 "vrepoc": "07.04.2023 01:35:00",
			 "vrezav": "07.04.2023 02:25:00",
			 "id_s_grraz": "3",
			 "id_s_razlog": "21",
			 "man_tekst": "Od 23.03.23 07:45 na snazi planirani radovi broj: R-67/1 (AKZ metalnih nosača VN opreme u TRP 400kV od TR1 u  TS Beograd 8 Vasiljević Aleksandar). NA SNAZI;",
			 "id_s_nap": "",
			 "p2_traf_id": "",
			 "kor_uneo": "pi1100",
			 "ed_id": "1"}`,
			"",
			app.updatePiPiDDNIsklj,
			http.StatusBadRequest,
		},
		{
			"insertPiPiDDNIsklj valid",
			"PUT",
			`{"dat_smene": "07.04.2023",
			"id_s_mrc": "8",
			 "tip_man": "1",
			 "id_tipob": "4",
			 "ob_id": "105",
			 "trafo_id": "",
			 "vrepoc": "07.04.2023 01:35:00",
			 "vrezav": "07.04.2023 02:25:00",
			 "id_s_grraz": "3",
			 "id_s_razlog": "21",
			 "man_tekst": "Od 23.03.23 07:45 na snazi planirani radovi broj: R-67/1 (AKZ metalnih nosača VN opreme u TRP 400kV od TR1 u  TS Beograd 8 Vasiljević Aleksandar). NA SNAZI;",
			 "id_s_nap": "",
			 "p2_traf_id": "",
			 "kor_uneo": "pi1100",
			 "ed_id": "1"}`,
			"",
			app.insertPiPiDDNIsklj,
			http.StatusAccepted,
		},
		{
			"insertPiPiDDNIsklj invalid",
			"PUT",
			`{"dat_smene": "07.04.2023",
			"id_s_mrc": "",
			 "tip_man": "1",
			 "id_tipob": "4",
			 "ob_id": "105",
			 "trafo_id": "",
			 "vrepoc": "07.04.2023 01:35:00",
			 "vrezav": "07.04.2023 02:25:00",
			 "id_s_grraz": "3",
			 "id_s_razlog": "21",
			 "man_tekst": "Od 23.03.23 07:45 na snazi planirani radovi broj: R-67/1 (AKZ metalnih nosača VN opreme u TRP 400kV od TR1 u  TS Beograd 8 Vasiljević Aleksandar). NA SNAZI;",
			 "id_s_nap": "",
			 "p2_traf_id": "",
			 "kor_uneo": "pi1100",
			 "ed_id": "1"}`,
			"",
			app.insertPiPiDDNIsklj,
			http.StatusBadRequest,
		},
		{
			"insertPiPiDDNIsklj invalid json",
			"PUT",
			`{"dat_smene": "07.04.2023",
			id_s_mrc "8",
			 "tip_man": "1",
			 "id_tipob": "4",
			 "ob_id": "105",
			 "trafo_id": "",
			 "vrepoc": "07.04.2023 01:35:00",
			 "vrezav": "07.04.2023 02:25:00",
			 "id_s_grraz": "3",
			 "id_s_razlog": "21",
			 "man_tekst": "Od 23.03.23 07:45 na snazi planirani radovi broj: R-67/1 (AKZ metalnih nosača VN opreme u TRP 400kV od TR1 u  TS Beograd 8 Vasiljević Aleksandar). NA SNAZI;",
			 "id_s_nap": "",
			 "p2_traf_id": "",
			 "kor_uneo": "pi1100",
			 "ed_id": "1"}`,
			"",
			app.insertPiPiDDNIsklj,
			http.StatusBadRequest,
		},
	}

	for _, e := range tests {
		var req *http.Request
		if e.json == "" {
			req, _ = http.NewRequest(e.method, "/", nil)
		} else {
			req, _ = http.NewRequest(e.method, "/", strings.NewReader(e.json))
		}

		if e.paramID != "" {
			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("id", e.paramID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(e.handler)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatus {
			t.Errorf("%s: wrong status returned; expected %d but got %d", e.name, e.expectedStatus, rr.Code)
		}
	}
}
