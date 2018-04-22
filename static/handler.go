package static

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/Alma-media/go-middleware"
	"github.com/Alma-media/restful"
	mwr "github.com/Alma-media/restful/middleware"
	"github.com/gorilla/mux"
)

// Module represents single static restful module.
type Handler struct {
	*mux.Router
	// it is strongly required to keep the order of modules as it was defined
	// since it can influence the priorities of application routing
	aliases []string
	modules []Module
}

// Use registers the module with provided alias and returns
func (h *Handler) Use(alias string, module Module) error {
	for i := 0; i < len(h.aliases); i++ {
		if h.aliases[i] == alias || h.modules[i] == module {
			return fmt.Errorf("alias/module already in use %q", alias)
		}
	}
	h.aliases = append(h.aliases, alias)
	h.modules = append(h.modules, module)
	return nil
}

// TODO: subroutes, use internal mux serve HTTP, register API on top level

// Init static Module. TODO: use ServeHTTP and hide router
func (h *Handler) Init() error {
	mux := h.Router
	// build routes for registered controllers
	for index, modulePath := range h.aliases {
		h.modules[index].Controllers(func(controllerPath string, resource restful.Controller) bool {
			basePath := []string{"/", modulePath, controllerPath}
			// list of available methods for current resource (required for OPTIONS request)
			var allowedMethods = make(map[string]struct{})
			// [GET] plural
			if controller, ok := resource.(PluralGetter); ok {
				mux.Handle(
					path.Join(basePath...),
					// apply default middleware (no need to close the body with mw.BodyClose)
					mw.New(mw.PanicRecover(mw.PanicHandler), mwr.Codec, mwr.GorillaParams).
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
				mux.Handle(
					path.Join(append(basePath, "{pk}")...),
					// apply default middleware (no need to close the body with mw.BodyClose)
					mw.New(mw.PanicRecover(mw.PanicHandler), mwr.Codec, mwr.GorillaParams).
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
				mux.Handle(
					path.Join(basePath...),
					// apply default middleware
					mw.New(mw.PanicRecover(mw.PanicHandler), mwr.Codec, mw.BodyClose, mwr.GorillaParams).
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
				mux.Handle(
					path.Join(append(basePath, "{pk}")...),
					// apply default middleware
					mw.New(mw.PanicRecover(mw.PanicHandler), mwr.Codec, mw.BodyClose, mwr.GorillaParams).
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
	}
	return nil
}
