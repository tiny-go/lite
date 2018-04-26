package user

import (
	"context"
	"log"

	"github.com/Alma-media/restful"
	"github.com/Alma-media/restful/examples/config"
	mw "github.com/tiny-go/middleware"
)

var (
	// compile-time type check (Controller should implement both interfaces)
	_ static.PluralPoster  = &Controller{}
	_ static.PluralPatcher = &Controller{}
)

// Controller is responsible for user AUTH operations.
type Controller struct {
	*mw.BaseController
	Config *config.Config    `inject:"t"`
	Users  map[string]string `inject:"t"`
}

// Init user controller (TODO: add middleware for available methods).
func (c *Controller) Init() error {
	// check that config was passed
	log.Println(c.Config, c.Users)

	return nil
}

// PostAll handles user login request.
func (c *Controller) PostAll(ctx context.Context, cf func(interface{}) error) (interface{}, error) {
	auth := new(Auth)
	if err := cf(auth); err != nil {
		return nil, err
	}
	return auth, auth.Login(c.Users)
}

// PatchAll handles user token refresh request.
func (c *Controller) PatchAll(ctx context.Context, cf func(interface{}) error) (interface{}, error) {
	auth := new(Auth)
	if err := cf(auth); err != nil {
		return nil, err
	}
	return auth, auth.RefreshToken()
}
