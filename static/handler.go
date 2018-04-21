package static

import (
	"log"
	"net/http"
	"path"

	"github.com/Alma-media/restful"
	mwr "github.com/Alma-media/restful/middleware"

	"github.com/Alma-media/go-middleware"
	"github.com/gorilla/mux"
)

// Module represents single static restful module.
type Module struct {
	Name      string
	Resources map[string]restful.Controller
}

// Init static Module.
func (m *Module) Init(mux *mux.Router) error {
	// TODO: deregister existing module routes before building new ones
	for name, resource := range m.Resources {
		basePath := []string{"/", m.Name, name}
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
	}
	return nil
}
