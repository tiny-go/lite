package lite

import (
	"context"
	"net/url"
)

// SingleGetter should be able to provide a single model by primary key(s)
type SingleGetter interface {
	Controller
	Get(ctx context.Context, pk string) (interface{}, error)
}

// PluralGetter should be able to provide a list of available models.
type PluralGetter interface {
	Controller
	GetAll(ctx context.Context, params url.Values) (interface{}, error)
}

// SinglePoster should be able to store a single model to the storage.
type SinglePoster interface {
	Controller
	Post(ctx context.Context, f func(v interface{}) error) (interface{}, error)
}

// PluralPoster should be able to store a list of model to the storage.
type PluralPoster interface {
	Controller
	PostAll(ctx context.Context, f func(v interface{}) error) (interface{}, error)
}

// SinglePatcher should be able to patch a single model by primary key(s).
type SinglePatcher interface {
	Controller
	Patch(ctx context.Context, pk string, f func(v interface{}) error) (interface{}, error)
}

// PluralPatcher should be able to patch a list of models.
type PluralPatcher interface {
	Controller
	PatchAll(ctx context.Context, params url.Values, f func(v interface{}) error) (interface{}, error)
}

// SinglePutter should be able to update a single model by primary key(s).
type SinglePutter interface {
	Controller
	Put(ctx context.Context, pk string, f func(v interface{}) error) (interface{}, error)
}

// PluralPutter should be able to update a list of models.
type PluralPutter interface {
	Controller
	PutAll(ctx context.Context, params url.Values, f func(v interface{}) error) (interface{}, error)
}

// SingleDeleter should be able to delete a single model by primary key(s).
type SingleDeleter interface {
	Controller
	Delete(ctx context.Context, pk string) (interface{}, error)
}

// PluralDeleter should be able to delete a list of models.
type PluralDeleter interface {
	Controller
	DeleteAll(ctx context.Context, params url.Values) (interface{}, error)
}
