package user

import (
	"fmt"
	"net/http"

	"github.com/tiny-go/errors"
)

// Auth is a user model.
type Auth struct {
	Email    string `json:"email,omitempty" xml:"Email,omitempty"`
	Password string `json:"password,omitempty" xml:"Password,omitempty"`
	Token    string `json:"token,omitempty" xml:"Token,omitempty"`
	Refresh  string `json:"refresh,omitempty" xml:"Refresh,omitempty"`
}

// Login with email and password.
func (a *Auth) Login(registry map[string]string) error {
	// simple auth logic (should be implemented with database instead of map)
	password, ok := registry[a.Email]
	if !ok || password != a.Password {
		return errors.NewStatusError(http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
	}
	// cleanup email/password fields
	a.Email, a.Password = "", ""

	return nil
}

// RefreshToken returns new token and refresh tokens.
func (a *Auth) RefreshToken() error {
	// TODO: add some logic
	return nil
}
