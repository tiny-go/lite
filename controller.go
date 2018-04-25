package static

import (
	"github.com/tiny-go/middleware"
)

// Controller interface is a bare minimal controller.
type Controller interface{
	 mw.Controller
	 Init() error
}

// BaseController ...
type BaseController = mw.BaseController
