package middleware

import (
	"context"
	"net/http"

	"github.com/notnmeyer/dngl/internal/envhelper"
)

func ContextInjector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "env", envhelper.New())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
