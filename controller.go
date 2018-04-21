package restful

import "github.com/Alma-media/go-middleware"

// Middleware represents HTTP middleware.
type Middleware = mw.Middleware

// Controller describes data controller.
type Controller interface {
	// AddMiddleware should make the provided chain available by HTTP method.
	AddMiddleware(method string, chain ...Middleware) Controller
	// Middleware should return the list of middleware functions registered for provided method.
	Middleware(method string) Middleware
}
