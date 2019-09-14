package exec

import (
	"net/http"

	"github.com/tiny-go/iauth"
	"github.com/tiny-go/lite"
	mw "github.com/tiny-go/middleware"
)

// add auth module to global module registry. Preferably do it with init func: to
// have only one instance of the module (keep it as a rule using global registry).
func init() {
	clients := make(iauth.MapClientStorage)
	clients.Register(iauth.Client{
		Key:    "5e0cf2d5-8eaa-4a58-b960-a409178f7e45",
		Secret: "passphrase",
	})

	cache := make(iauth.MapCache)
	authMW := &iauth.Manager{Cache: cache, ClientStorage: clients}

	module := lite.NewBaseModule()
	controller := &Controller{BaseController: mw.NewBaseController()}

	controller.AddMiddleware(http.MethodOptions, authMW.Preflight())

	module.Register("", controller)
	lite.Register("exec", module)
}
