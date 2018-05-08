package lite

import (
	"context"
	"errors"

	"github.com/tiny-go/middleware"
)

var (
	_ SingleGetter  = &mockController{}
	_ PluralGetter  = &mockController{}
	_ SinglePoster  = &mockController{}
	_ PluralPoster  = &mockController{}
	_ SinglePatcher = &mockController{}
	_ PluralPatcher = &mockController{}
	_ SinglePutter  = &mockController{}
	_ PluralPutter  = &mockController{}
	_ SingleDeleter = &mockController{}
	_ PluralDeleter = &mockController{}
)

type mockController struct {
	mw.Controller
	ShouldFail bool
}

func newPassController() *mockController {
	return &mockController{mw.NewBaseController(), false}
}

func newFailController() *mockController {
	return &mockController{mw.NewBaseController(), true}
}

func (c *mockController) Init() error { return nil }

func (c *mockController) Get(_ context.Context, pk string) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("single GET error")
	}
	return pk, nil
}

func (c *mockController) GetAll(context.Context) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("plural GET error")
	}
	return map[string]interface{}{"foo": "bar"}, nil
}

func (c *mockController) Post(_ context.Context, pk string, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("single POST error")
	}
	data := map[string]interface{}{"pk": pk}
	return data, f(&data)
}

func (c *mockController) PostAll(_ context.Context, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("plural POST error")
	}
	data := make(map[string]interface{})
	return data, f(&data)
}

func (c *mockController) Patch(_ context.Context, pk string, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("single PATCH error")
	}
	data := map[string]interface{}{"pk": pk}
	return data, f(&data)
}

func (c *mockController) PatchAll(_ context.Context, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("plural PATCH error")
	}
	data := make(map[string]interface{})
	return data, f(&data)
}

func (c *mockController) Put(_ context.Context, pk string, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("single PUT error")
	}
	data := map[string]interface{}{"pk": pk}
	return data, f(&data)
}

func (c *mockController) PutAll(_ context.Context, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("plural PUT error")
	}
	data := make(map[string]interface{})
	return data, f(&data)
}

func (c *mockController) Delete(_ context.Context, pk string) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("single DELETE error")
	}
	return pk, nil
}

func (c *mockController) DeleteAll(_ context.Context, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("plural DELETE error")
	}
	data := make(map[string]interface{})
	return data, f(&data)
}
