package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodPost, "/signin", app.Signin)
	router.HandlerFunc(http.MethodGet, "/signals/dv_didf", app.getDvDidf)
	router.HandlerFunc(http.MethodGet, "/signals/tr_diff", app.getDiffTr)
	router.HandlerFunc(http.MethodGet, "/signals/tr_dis_res", app.getDisTrRes)
	router.HandlerFunc(http.MethodGet, "/signals/sp_dis_diff", app.getDisDiffSp)
	router.HandlerFunc(http.MethodGet, "/signals/malfunction_in", app.getMalfunctionIn)
	router.HandlerFunc(http.MethodGet, "/signals/apu", app.getAPU)
	router.HandlerFunc(http.MethodGet, "/signals/dv_oc", app.getOCDV)

	router.POST("/dv_didf", app.wrap(secure.ThenFunc(app.getDvDidf)))
	return app.enableCORS(router)

}
