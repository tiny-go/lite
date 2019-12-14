package user

import (
	"context"
	"net/url"

	"github.com/tiny-go/lite"
	"github.com/tiny-go/lite/examples/ums/config"
	mw "github.com/tiny-go/middleware"
)

var (
	// compile-time type check (Controller should implement both interfaces)
	_ lite.PluralPoster  = &Controller{}
	_ lite.PluralPatcher = &Controller{}
)

// Controller is responsible for user AUTH operations.
type Controller struct {
	// inherit BaseController functionality (related to middleware)
	*mw.BaseController
	// controller dependencies
	Config *config.Config    `inject:"t"`
	Users  map[string]string `inject:"t"`
}

// Init user controller (TODO: add middleware for available methods).
func (c *Controller) Init() error { return nil }

// PostAll handles user login request.
func (c *Controller) PostAll(_ context.Context, cf func(interface{}) error) (interface{}, error) {
	auth := new(Auth)
	if err := cf(auth); err != nil {
		return nil, err
	}
	return auth, auth.Login(c.Config, c.Users)
}

// PatchAll handles user token refresh request.
func (c *Controller) PatchAll(_ context.Context, _ url.Values, cf func(interface{}) error) (interface{}, error) {
	auth := new(Auth)
	if err := cf(auth); err != nil {
		return nil, err
	}
	return auth, auth.RefreshToken()
}
