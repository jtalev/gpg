package main

import (
	"context"
	"gpg/portal/internal/handle"
	"gpg/portal/internal/localdb"
	"net/http"
)


func addRoutes(mux *http.ServeMux, ctx context.Context, db localdb.Db) {
	mux.Handle("/", handle.ServeIndex(ctx, db))
}