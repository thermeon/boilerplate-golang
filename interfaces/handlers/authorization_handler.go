package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/thermeon/boilerplate-golang/infrastructures/server/operations"
	"github.com/thermeon/boilerplate-golang/log"
)

func GetAuthorizationTransactionHandlerFunc() operations.AuthorizationTransactionHandlerFunc {
	return func(params operations.AuthorizationTransactionParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		log.Debug(ctx, "Hitting Authorization Transaction Handler. Principles: %s, Params: %s", principal, params)
		return nil
	}
}
