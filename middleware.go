package static

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type paramsKey struct{}

// Params is a key/value map containing the request paramters from the URI.
type Params map[string]string

// GorillaParams extracts URI params from the request (when using gorolla/mux).
func GorillaParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		// convert gorilla Params to internal representation
		ps := make(Params, len(vars))
		for key, value := range vars {
			ps[key] = value
		}
		r = r.WithContext(context.WithValue(r.Context(), paramsKey{}, ps))
		// call the next handler
		next.ServeHTTP(w, r)
	})
}

// ParamsFromContext pulls the URL parameters from a request context,
// or returns nil if none are present.
func ParamsFromContext(ctx context.Context) Params {
	p, _ := ctx.Value(paramsKey{}).(Params)
	return p
}
