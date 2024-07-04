package main

import (
	"context"
	"gpg/portal/internal/handle"
	"gpg/portal/internal/localdb"
	"net/http"
)


func addRoutes(mux *http.ServeMux, ctx context.Context, db localdb.Db) {
	fileServer := http.FileServer(http.Dir("../../web/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
    mux.Handle("/", handle.ServeIndex(ctx, db))
}