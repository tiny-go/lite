package auth

import (
	"github.com/Alma-media/restful"
	"github.com/Alma-media/restful/examples/auth/user"
	mw "github.com/tiny-go/middleware"
)

// add auth module to global module registry. Preferably do it with init func: to
// have only one instance of the module (keep it as a rule using global registry).
func init() {
	module := static.NewBaseModule()
	module.Register("", &user.Controller{BaseController: mw.NewBaseController()})
	static.Register("auth", module)
}
