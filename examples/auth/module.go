package auth

import (
	"github.com/Alma-media/restful"
	"github.com/Alma-media/restful/examples/auth/user"
	mw "github.com/tiny-go/middleware"
)

// add auth module to global module registry
func init() {
	module := static.NewBaseModule()
	module.Register("", &user.Controller{BaseController: mw.NewBaseController()})
	static.Register("auth", module)
}
