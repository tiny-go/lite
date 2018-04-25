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
	// modules is a local registry which is needed for alias/module unique check
	modules map[string]Module
}

// NewHandler creates new static handler.
func NewHandler() *Handler {
	return &Handler{Router: mux.NewRouter(), modules: make(map[string]Module)}
}

// Use registers the module with provided alias.
func (h *Handler) Use(alias string, module Module) (err error) {
	for key, value := range h.modules {
		if key == alias || value == module {
			return fmt.Errorf("alias/module already in use %q", alias)
		}
	}

	module.Controllers(func(controllerPath string, resource Controller) bool {
		// init current controller first and if failed return false to exit the loop
		if err = resource.Init(); err != nil {
			return false
		}
		basePath := []string{"/", alias, controllerPath}
		// list of available methods for current resource (required for OPTIONS request)
		var allowedMethods = make(map[string]struct{})
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
		// [PATCH] plural
		if controller, ok := resource.(PluralPatcher); ok {
			h.Router.Handle(
				path.Join(basePath...),
				// apply default middleware
				mw.New(mw.PanicRecover(errors.Send), mw.Codec(driver.Global()), mw.BodyClose, GorillaParams).
					// extract custom (user defined) middleware for HTTP method PATCH
					Use(controller.Middleware(http.MethodPatch)).
					// set final handler
					Then(patchPlural(controller)),
			).Methods(http.MethodPatch)
			// add PATCH method to OPTIONS list
			allowedMethods[http.MethodPatch] = struct{}{}

			log.Printf("[PATCH] %s\n", path.Join(basePath...))
		}
		// [PATCH] single
		if controller, ok := resource.(SinglePatcher); ok {
			h.Router.Handle(
				path.Join(append(basePath, "{pk}")...),
				// apply default middleware
				mw.New(mw.PanicRecover(errors.Send), mw.Codec(driver.Global()), mw.BodyClose, GorillaParams).
					// extract custom (user defined) middleware for HTTP method PATCH
					Use(controller.Middleware(http.MethodPatch)).
					// set final handler
					Then(patchSingle(controller)),
			).Methods(http.MethodPatch)
			// add POST method to OPTIONS list
			allowedMethods[http.MethodPatch] = struct{}{}

			log.Printf("[PATCH] %s\n", path.Join(append(basePath, "{pk}")...))
		}
		// [PUT] plural
		if controller, ok := resource.(PluralPutter); ok {
			h.Router.Handle(
				path.Join(basePath...),
				// apply default middleware
				mw.New(mw.PanicRecover(errors.Send), mw.Codec(driver.Global()), mw.BodyClose, GorillaParams).
					// extract custom (user defined) middleware for HTTP method PUT
					Use(controller.Middleware(http.MethodPut)).
					// set final handler
					Then(putPlural(controller)),
			).Methods(http.MethodPut)
			// add PATCH method to OPTIONS list
			allowedMethods[http.MethodPut] = struct{}{}

			log.Printf("[PUT] %s\n", path.Join(basePath...))
		}
		// [PUT] single
		if controller, ok := resource.(SinglePutter); ok {
			h.Router.Handle(
				path.Join(append(basePath, "{pk}")...),
				// apply default middleware
				mw.New(mw.PanicRecover(errors.Send), mw.Codec(driver.Global()), mw.BodyClose, GorillaParams).
					// extract custom (user defined) middleware for HTTP method PUT
					Use(controller.Middleware(http.MethodPut)).
					// set final handler
					Then(putSingle(controller)),
			).Methods(http.MethodPut)
			// add POST method to OPTIONS list
			allowedMethods[http.MethodPut] = struct{}{}

			log.Printf("[PUT] %s\n", path.Join(append(basePath, "{pk}")...))
		}

		// TODO: implement for the rest of HTTP methods

		return true
	})
	// store alias and module to local registry in order to avoid duplicates
	h.modules[alias] = module

	return err
}
