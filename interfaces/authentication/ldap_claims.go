package authentication

import "github.com/golang-jwt/jwt"

type LdapClaims struct {
	Groups []string `json:"groups,omitempty"`
	jwt.StandardClaims
}
