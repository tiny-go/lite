package restful

import "github.com/Alma-media/go-middleware"

// Controller represents simple HTTP controller containing middleware for each method.
type Controller interface {
	// AddMiddleware should make the provided chain available by HTTP method.
	AddMiddleware(method string, chain ...mw.Middleware) Controller
	// Middleware should return the list of middleware functions registered for provided method.
	Middleware(method string) mw.Middleware
}
