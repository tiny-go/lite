package static

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/tiny-go/codec/driver"
	"github.com/tiny-go/errors"
	"github.com/tiny-go/middleware"
)

// Handler combines all static modules (with their controllers) to a single API.
type Handler struct {
	*mux.Router
	// it is strongly required to keep the order of modules as it was defined
	// since it can influence the priorities of application routing
	modules map[string]Module
}

// NewHandler creates new static handler.
func NewHandler() *Handler {
	return &Handler{Router: mux.NewRouter(), modules: make(map[string]Module)}
}

// Use registers the module with provided alias and returns
func (h *Handler) Use(alias string, module Module) error {
	for key, value := range h.modules {
		if key == alias || value == module {
			return fmt.Errorf("alias/module already in use %q", alias)
		}
	}

	module.Controllers(func(controllerPath string, resource Controller) bool {
		basePath := []string{"/", alias, controllerPath}
		// list of available methods for current resource (required for OPTIONS request)
		var allowedMethods = make(map[string]struct{})

		resource.Init()
		// [GET] plural
		if controller, ok := resource.(PluralGetter); ok {
			h.Router.Handle(
				path.Join(basePath...),
				// apply default middleware (no need to close the body with mw.BodyClose)
				mw.New(mw.PanicRecover(errors.Send), mw.Codec(driver.Global()), GorillaParams).
					// extract custom (user defined) middleware for HTTP method GET
					Use(controller.Middleware(http.MethodGet)).
					// set final handler
					Then(getPlural(controller)),
			).Methods(http.MethodGet)
			// add GET method to OPTIONS list
			allowedMethods[http.MethodGet] = struct{}{}

			log.Printf("[GET] %s\n", path.Join(basePath...))
		}
		// [GET] single
		if controller, ok := resource.(SingleGetter); ok {
			h.Router.Handle(
				path.Join(append(basePath, "{pk}")...),
				// apply default middleware (no need to close the body with mw.BodyClose)
				mw.New(mw.PanicRecover(errors.Send), mw.Codec(driver.Global()), GorillaParams).
					// extract custom (user defined) middleware for HTTP method GET
					Use(controller.Middleware(http.MethodGet)).
					// set final handler
					Then(getSingle(controller)),
			).Methods(http.MethodGet)
			// add GET method to OPTIONS list
			allowedMethods[http.MethodGet] = struct{}{}

			log.Printf("[GET] %s\n", path.Join(append(basePath, "{pk}")...))
		}
		// [POST] plural
		if controller, ok := resource.(PluralPoster); ok {
			h.Router.Handle(
				path.Join(basePath...),
				// apply default middleware
				mw.New(mw.PanicRecover(errors.Send), mw.Codec(driver.Global()), mw.BodyClose, GorillaParams).
					// extract custom (user defined) middleware for HTTP method POST
					Use(controller.Middleware(http.MethodPost)).
					// set final handler
					Then(postPlural(controller)),
			).Methods(http.MethodPost)
			// add POST method to OPTIONS list
			allowedMethods[http.MethodPost] = struct{}{}

			log.Printf("[POST] %s\n", path.Join(basePath...))
		}
		// [POST] single
		if controller, ok := resource.(SinglePoster); ok {
			h.Router.Handle(
				path.Join(append(basePath, "{pk}")...),
				// apply default middleware
				mw.New(mw.PanicRecover(errors.Send), mw.Codec(driver.Global()), mw.BodyClose, GorillaParams).
					// extract custom (user defined) middleware for HTTP method POST
					Use(controller.Middleware(http.MethodPost)).
					// set final handler
					Then(postSingle(controller)),
			).Methods(http.MethodPost)
			// add POST method to OPTIONS list
			allowedMethods[http.MethodPost] = struct{}{}

			log.Printf("[POST] %s\n", path.Join(append(basePath, "{pk}")...))
		}
		// TODO: implement for the rest of HTTP methods

		return true
	})

	h.modules[alias] = module

	return nil
}
