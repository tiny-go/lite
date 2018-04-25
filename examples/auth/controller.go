package auth

import (
	"context"

	"github.com/Alma-media/restful"
)

var _ static.PluralPoster = &Controller{}

type Controller struct {
	*static.BaseController
}

func (c *Controller) PostAll(ctx context.Context, cf func(interface{}) error) (interface{}, error) {
	model := new(Model)
	if err := cf(model); err != nil {
		return nil, err
	}

	return model, nil
}
