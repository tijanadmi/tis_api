package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_routes(t *testing.T) {
	var registered = []struct {
		route  string
		method string
	}{
		{"/authenticate", "POST"},
		{"/refresh", "GET"},
		{"/logout", "GET"},
		{"/status", "GET"},
		{"/dwh/weather/{year}", "GET"},
		{"/dwh/weatherforecast", "GET"},
		{"/dwh/weatherhistory/{year}", "GET"},
		{"/dwh/dozvole1/{year}", "GET"},
		{"/dwh/dozvole23/{year}", "GET"},
		{"/dwh/zahtevi1/{year}", "GET"},
		{"/dwh/zahtevi23/{year}", "GET"},
		{"/dwh/ispadi/{year}", "GET"},
		{"/dwh/iskljucenja/{year}", "GET"},
		{"/dwh/planovi/{year}", "GET"},
		{"/admin/signals/{id}", "GET"},
		{"/admin/signals/dvdidf", "GET"},
		{"/admin/signals/trdiff", "GET"},
		{"/admin/signals/trdisres", "GET"},
		{"/admin/signals/spdisdiff", "GET"},
		{"/admin/signals/malfunctionin", "GET"},
		{"/admin/signals/malfunctionin/{id}", "GET"},
		{"/admin/signals/apu", "GET"},
		{"/admin/signals/apu/{id}", "GET"},
		{"/admin/signals/dvoc", "GET"},
		{"/admin/signals/tr12oc", "GET"},
		{"/admin/signals/trresoc", "GET"},
		{"/admin/signals/spocsc", "GET"},
		{"/admin/signals/dvearthfaultoc", "GET"},
		{"/admin/signals/trearthfaultoc", "GET"},
		{"/admin/signals/spearthfaultoc", "GET"},
		{"/admin/signals/direarthfaultoc", "GET"},
		{"/admin/signals/tpsendrcvd", "GET"},
		{"/admin/signals/circuitbreaker", "GET"},
		{"/admin/signals/bbpbftrip", "GET"},
		{"/admin/signals/nonelectrical", "GET"},
		{"/admin/signals/bbpbbtrip", "GET"},
		{"/admin/signals/bftrip", "GET"},
		{"/admin/groupsofcauses", "GET"},
		{"/admin/groupsofcauses/{id}", "GET"},
		{"/admin/causes", "GET"},
		{"/admin/causes/{id}", "GET"},
		{"/admin/groupsofreasons", "GET"},
		{"/admin/groupsofreasons/{id}", "GET"},
		{"/admin/reasons", "GET"},
		{"/admin/reasons/{id}", "GET"},
		{"/admin/weatherconditions", "GET"},
		{"/admin/weatherconditions/{id}", "GET"},
		{"/admin/categoriesofevents", "GET"},
		{"/admin/categoriesofevents/{id}", "GET"},
		{"/admin/overheadlines", "GET"},
		{"/admin/powercables", "GET"},
		{"/admin/substations", "GET"},
		{"/admin/feeders", "GET"},
		{"/admin/protectiondevices", "GET"},
		{"/admin/powertransformers", "GET"},
		{"/admin/disconnectors", "GET"},
		{"/admin/workpermissions", "GET"},
		{"/admin/workineenetwork", "GET"},
		{"/admin/workpermissionsall", "GET"},
		{"/admin/request1gr", "GET"},
		{"/admin/request2gr", "GET"},
		{"/admin/unopenedpermitsforday/{org}/{day}", "GET"},
		{"/admin/unfinishedevents/ndc", "GET"},
		{"/admin/unfinishedevents/{id}", "GET"},
		{"/admin/unfinishedevents/{id}", "PATCH"},
		{"/admin/interruptionofdelivery/ndc", "GET"},
		{"/admin/interruptionofdelivery/{id}", "GET"},
		{"/admin/interruptionofdelivery/0", "PUT"},
		{"/admin/interruptionofdelivery/{id}", "PATCH"},
		{"/admin/interruptionofdelivery/{id}", "DELETE"},
		{"/admin/pipiddn", "GET"},
		{"/admin/pipiddn/{id}", "GET"},
		{"/admin/pipiddn/operation/0", "PUT"},
		{"/admin/pipiddn/outage/0", "PUT"},
		{"/admin/pipiddn/operation/{id}", "PATCH"},
		{"/admin/pipiddn/outage/{id}", "PATCH"},
		{"/admin/pipiddn/{id}", "DELETE"},
		{"/admin/unbalancestraders", "GET"},
	}

	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, route := range registered {
		// check to see if the route exists
		if !routeExists(route.route, route.method, chiRoutes) {
			t.Errorf("route %s is not registered", route.route)
		}
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})

	return found
}
