package user

import (
	"context"

	"github.com/Alma-media/restful"
	mw "github.com/tiny-go/middleware"
)

var (
	// compile-time type check (Controller should implement both interfaces)
	_ static.PluralPoster  = &Controller{}
	_ static.PluralPatcher = &Controller{}
)

// Controller is responsible for user AUTH operations.
type Controller struct{ *mw.BaseController }

// Init user controller (TODO: add middleware for available methods).
func (c *Controller) Init() error { return nil }

// PostAll handles user login request.
func (c *Controller) PostAll(ctx context.Context, cf func(interface{}) error) (interface{}, error) {
	model := new(User)
	if err := cf(model); err != nil {
		return nil, err
	}
	return model, model.Login()
}

// PatchAll handles user token refresh request.
func (c *Controller) PatchAll(ctx context.Context, cf func(interface{}) error) (interface{}, error) {
	model := new(User)
	if err := cf(model); err != nil {
		return nil, err
	}
	return model, model.RefreshToken()
}
