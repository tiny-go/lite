package auth

import (
	"github.com/tiny-go/lite"
	"github.com/tiny-go/lite/example/auth/user"
	mw "github.com/tiny-go/middleware"
)

// add auth module to global module registry. Preferably do it with init func: to
// have only one instance of the module (keep it as a rule using global registry).
func init() {
	module := lite.NewBaseModule()
	module.Register("", &user.Controller{BaseController: mw.NewBaseController()})
	lite.Register("auth", module)
}
