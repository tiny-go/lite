package auth

import (
	"context"
	"net/http"

	"github.com/Alma-media/restful"
	"github.com/tiny-go/codec/driver"
	mw "github.com/tiny-go/middleware"
)

var _ static.PluralPoster = &UserController{}

// UserController is responsible for user AUTH operations.
type UserController struct{ *mw.BaseController }

// Init user controller (add middleware for available methods).
func (c *UserController) Init() error {
	c.Middleware(http.MethodPost).
		Use(mw.Codec(driver.Global()))
	return nil
}

// PostAll is used to login.
func (c *UserController) PostAll(ctx context.Context, cf func(interface{}) error) (interface{}, error) {
	model := new(UserModel)
	if err := cf(model); err != nil {
		return nil, err
	}
	// TODO: for now it does nothing - add some logic
	return model, nil
}
