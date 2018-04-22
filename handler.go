package restful

import "net/http"

// Handler describes a single application module.
type Handler interface {
	http.Handler
	// Init should initialize a handler buildig routes for all registered modules.
	// Call this func before building routes for your application .
	Init() error
}
