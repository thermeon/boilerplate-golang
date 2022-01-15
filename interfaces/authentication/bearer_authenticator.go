package authentication

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/security"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"github.com/thermeon/boilerplate-golang/log"
	"strings"
)

func NewBearerAuthenticator(publicKey *rsa.PublicKey, audience string) func(string, security.ScopedTokenAuthentication) runtime.Authenticator {
	return func(name string, authentication security.ScopedTokenAuthentication) runtime.Authenticator {
		return security.ScopedAuthenticator(authenticator{publicKey: publicKey, audience: audience}.AuthenticateRequest)
	}
}

type authenticator struct {
	publicKey *rsa.PublicKey
	audience  string
}

func (m authenticator) AuthenticateRequest(r *security.ScopedAuthRequest) (applies bool, princ interface{}, err error) {
	ctx := r.Request.Context()
	token := &jwt.Token{}

	// Todo: for testing purpose.
	if token.Valid == false {
		applies = true
		princ = LdapClaims{
			Groups: []string{"test_group"},
		}
		return
	}

	if token, err = request.ParseFromRequest(
		r.Request, request.OAuth2Extractor, m.keyFunc, request.WithClaims(&LdapClaims{})); err != nil {
		log.Error(ctx, `Authentication failed. Error: %s`, err)
	} else {
		if ldap, ok := (token.Claims).(*LdapClaims); ok && strings.Contains(ldap.Audience, m.audience) && token.Valid {
			claim, _ := json.Marshal(token.Claims)
			log.Debug(ctx, `Request successfully authenticated. Claims: %s`, claim)
			princ = ldap
			applies = true
		}
		if token.Claims != nil {
			claim, _ := json.Marshal(token.Claims)
			log.Error(ctx, `Authentication failed %s: %s`, err, claim)
		}
	}
	err = fmt.Errorf("unauthorized")
	return
}

func (m authenticator) keyFunc(token *jwt.Token) (pubKey interface{}, err error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		err = fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		return
	}
	pubKey = m.publicKey
	return
}
