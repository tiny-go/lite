package static

import (
	"github.com/Alma-media/go-middleware"
	"github.com/Alma-media/restful"
)

// BaseController provides ability to register/unregister middleware for HTTP methods.
// This is a basic Controller implementation.
type BaseController struct {
	middleware map[string]mw.Middleware
}

// NewBaseController is a constructor func for a basic HTTP controller.
func NewBaseController() *BaseController {
	return &BaseController{
		middleware: make(map[string]mw.Middleware),
	}
}

// AddMiddleware adds middleware funcs to existing ones for provided HTTP method.
func (bc *BaseController) AddMiddleware(method string, chain ...mw.Middleware) restful.Controller {
	if _, ok := bc.middleware[method]; !ok {
		// create new middleware
		bc.middleware[method] = mw.New(chain...)
	} else {
		// upgrade existing middleware and replace the old one
		bc.middleware[method] = bc.middleware[method].Use(chain...)
	}
	// return itself in order to use func in a chain
	return bc
}

// Middleware returns middleware func registered for provided method or an empty list.
func (bc *BaseController) Middleware(method string) mw.Middleware {
	if mw, ok := bc.middleware[method]; ok {
		return mw
	}
	return mw.New()
}
