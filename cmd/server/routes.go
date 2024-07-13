package main

import (
	"context"
	"gpg/portal/internal/database"
	"gpg/portal/internal/handle"
	"net/http"
)


func addRoutes(mux *http.ServeMux, ctx context.Context, db *database.Db) {
	fileServer := http.FileServer(http.Dir("../../web/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
    
	// pages
	mux.Handle("/", handle.ServeLogin(ctx, db))
	mux.Handle("/dashboard", handle.ServeDashboard(ctx, db))
	mux.Handle("/admin", handle.ServeAdmin())

	// login status
	mux.Handle("/validate-login", handle.HandleValidateLogin(ctx, db))
	mux.Handle("/logout", handle.HandleLogout())
}