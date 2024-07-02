package main

import (
	"encoding/json"
	"net/http"
)

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello")
}