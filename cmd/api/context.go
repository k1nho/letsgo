package main

import (
	"context"
	"net/http"

	"github.com/k1nho/letsgo/internal/data"
)

type contextKey string

const userContentKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContentKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContentKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
