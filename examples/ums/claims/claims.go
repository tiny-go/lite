package claims

import jwt "github.com/dgrijalva/jwt-go"

// JWT claims.
type JWT struct {
	jwt.StandardClaims
	Email string
}
