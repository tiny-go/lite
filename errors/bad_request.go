package errors

import (
	"errors"
	"net/http"

	"github.com/Alma-media/restful"
)

var _ restful.Error = BadRequest{}

// BadRequest is a 400 HTTP error.
type BadRequest struct {
	error
}

// Status returns http staus code.
func (e BadRequest) Status() int {
	return http.StatusBadRequest
}

// NewBadRequest is a constructor func for 400 HTTP error.
func NewBadRequest(cause string) BadRequest {
	return BadRequest{errors.New(cause)}
}
