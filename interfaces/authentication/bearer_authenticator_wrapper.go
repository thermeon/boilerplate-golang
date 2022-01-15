package authentication

import (
	"context"
	"crypto/rsa"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/security"
	"github.com/golang-jwt/jwt"
	"github.com/thermeon/boilerplate-golang/configs"
	"github.com/thermeon/boilerplate-golang/log"
	"io/ioutil"
)

func NewBearerAuthenticatorWrapper(conf *configs.Config) func(string, security.ScopedTokenAuthentication) runtime.Authenticator {
	var err error
	ctx := context.Background()
	pubKey := &rsa.PublicKey{}
	if conf.RSAPublicKeyPath == "" {
		log.Error(ctx, "RSA key path missing from configuration")
	}
	if pubKey, err = getPublicKey(conf.RSAPublicKeyPath); err != nil {
		log.Error(ctx, "RSA key could not be loaded. Error: %s", err)
	}
	return NewBearerAuthenticator(pubKey, conf.SecurityAudience)
}

func getPublicKey(path string) (pubKey *rsa.PublicKey, err error) {
	var data []byte
	if data, err = ioutil.ReadFile(path); err != nil {
		log.Error(context.Background(), "Unable to load the RSA public key. %v", err.Error())
		return
	}
	return jwt.ParseRSAPublicKeyFromPEM(data)
}
