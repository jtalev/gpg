package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type route struct {
	url		string
	f 		Handler
}

var routes = []route {
	{"/timesheet", handler}, 
}

func InitRoutes(r *mux.Router) {
	for _, rr := range routes {
		r.HandleFunc(rr.url, rr.f)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello")
}