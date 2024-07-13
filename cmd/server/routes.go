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
    
	// pages
	mux.Handle("/", handle.ServeLogin(ctx, db))
	mux.Handle("/dashboard", handle.ServeDashboard(ctx, db))

	// login status
	mux.Handle("/validate-login", handle.HandleValidateLogin(ctx, db))
	mux.Handle("/logout", handle.HandleLogout(db))
}