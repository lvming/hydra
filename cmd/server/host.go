package server

import (
	"context"
	"net/http"
)

func hostMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := context.WithValue(r.Context(), "Host", r.Host)
	next(rw, r.WithContext(ctx))
}
