package user

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tiny-go/errors"
	"github.com/tiny-go/lite/example/claims"
	"github.com/tiny-go/lite/example/config"
)

// Auth is a user model.
type Auth struct {
	Email    string `json:"email,omitempty" xml:"Email,omitempty"`
	Password string `json:"password,omitempty" xml:"Password,omitempty"`
	Token    string `json:"token,omitempty" xml:"Token,omitempty"`
	Refresh  string `json:"refresh,omitempty" xml:"Refresh,omitempty"`
}

// Login with email and password.
func (a *Auth) Login(config *config.Config, registry map[string]string) (err error) {
	// simple auth logic (should be implemented with database instead of map)
	password, ok := registry[a.Email]
	if !ok || password != a.Password {
		return errors.NewStatusError(http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
	}
	// create claims object
	claims := &claims.JWT{Email: a.Email}
	claims.ExpiresAt = time.Now().Add(config.TokLifeTime).Unix()
	// create a token from claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// encrypt token
	if a.Token, err = token.SignedString([]byte(config.Secret)); err != nil {
		return err
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
