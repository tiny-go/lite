package auth

import (
	"github.com/Alma-media/restful"
	mw "github.com/tiny-go/middleware"
)

// add auth module to global module registry
func init() {
	module := static.NewBaseModule()
	module.Register("user", &UserController{mw.NewBaseController()})
	static.Register("auth", module)
}
