package exec

import (
	"github.com/tiny-go/lite"
	mw "github.com/tiny-go/middleware"
)

// Add module to global module registry. Preferably do it with init func: to
// have only one instance of the module (keep it as a rule using global registry).
func init() {
	module := lite.NewBaseModule()
	controller := &Controller{BaseController: mw.NewBaseController()}
	module.Register("", controller)
	lite.Register("exec", module)
}
