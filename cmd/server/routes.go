package main

import (
	"context"
	"gpg/portal/internal/handle"
	"gpg/portal/internal/user"
	"net/http"
)


func addRoutes(mux *http.ServeMux, ctx context.Context) {
	u := user.User{
		Id: 1,
		EmployeeId: 1234567,
		Username: "Sliddy",
		FirstName: "Josh",
		LastName: "Talev",
	}
	ctx = context.WithValue(ctx, "user_name", u.Username)
	mux.Handle("/", handle.ServeIndex(ctx))
}