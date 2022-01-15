// This file is safe to edit. Once it exists it will not be overwritten

package server

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/thermeon/boilerplate-golang/bootstrap"
	"github.com/thermeon/boilerplate-golang/domain/entities"
	"github.com/thermeon/boilerplate-golang/interfaces/handlers"
	customMiddleware "github.com/thermeon/boilerplate-golang/interfaces/middleware"
	"net/http"

	"github.com/thermeon/boilerplate-golang/infrastructures/server/operations"
)

//go:generate swagger generate server --target ../../../payments-service --name PaymentsServiceAPI --spec ../../docs/spec/swagger.yaml --server-package ./infrastructures/server/. --principal interface{}

func configureFlags(api *operations.PaymentsServiceAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PaymentsServiceAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Initialize the payments service.
	bootstrap.Init(api)

	api.UseSwaggerUI()
	api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.AuthorizationTransactionHandler = handlers.GetAuthorizationTransactionHandlerFunc()

	api.PaymentsAuthAuth = func(token string, scopes []string) (interface{}, error) {
		// return nil, errors.NotImplemented("oauth2 bearer auth (payments_auth) has not yet been implemented")
		return entities.Transaction{ID: uuid.New()}, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.AbortTransactionHandler == nil {
		api.AbortTransactionHandler = operations.AbortTransactionHandlerFunc(func(params operations.AbortTransactionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.AbortTransaction has not yet been implemented")
		})
	}
	if api.AuthorizationTransactionHandler == nil {
		api.AuthorizationTransactionHandler = operations.AuthorizationTransactionHandlerFunc(func(params operations.AuthorizationTransactionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.AuthorizationTransaction has not yet been implemented")
		})
	}
	if api.CaptureTransactionHandler == nil {
		api.CaptureTransactionHandler = operations.CaptureTransactionHandlerFunc(func(params operations.CaptureTransactionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.CaptureTransaction has not yet been implemented")
		})
	}
	if api.ChooseOptionHandler == nil {
		api.ChooseOptionHandler = operations.ChooseOptionHandlerFunc(func(params operations.ChooseOptionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.ChooseOption has not yet been implemented")
		})
	}
	if api.CreateLaneHandler == nil {
		api.CreateLaneHandler = operations.CreateLaneHandlerFunc(func(params operations.CreateLaneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateLane has not yet been implemented")
		})
	}
	if api.DeleteLaneHandler == nil {
		api.DeleteLaneHandler = operations.DeleteLaneHandlerFunc(func(params operations.DeleteLaneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteLane has not yet been implemented")
		})
	}
	if api.GetLaneHandler == nil {
		api.GetLaneHandler = operations.GetLaneHandlerFunc(func(params operations.GetLaneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetLane has not yet been implemented")
		})
	}
	if api.GetLanesHandler == nil {
		api.GetLanesHandler = operations.GetLanesHandlerFunc(func(params operations.GetLanesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetLanes has not yet been implemented")
		})
	}
	if api.RefundTransactionHandler == nil {
		api.RefundTransactionHandler = operations.RefundTransactionHandlerFunc(func(params operations.RefundTransactionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.RefundTransaction has not yet been implemented")
		})
	}
	if api.RetrieveTransactionHandler == nil {
		api.RetrieveTransactionHandler = operations.RetrieveTransactionHandlerFunc(func(params operations.RetrieveTransactionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.RetrieveTransaction has not yet been implemented")
		})
	}
	if api.RetrieveTransactionHistoryHandler == nil {
		api.RetrieveTransactionHistoryHandler = operations.RetrieveTransactionHistoryHandlerFunc(func(params operations.RetrieveTransactionHistoryParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.RetrieveTransactionHistory has not yet been implemented")
		})
	}
	if api.SaleTransactionHandler == nil {
		api.SaleTransactionHandler = operations.SaleTransactionHandlerFunc(func(params operations.SaleTransactionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.SaleTransaction has not yet been implemented")
		})
	}
	if api.TokenizeCardHandler == nil {
		api.TokenizeCardHandler = operations.TokenizeCardHandlerFunc(func(params operations.TokenizeCardParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.TokenizeCard has not yet been implemented")
		})
	}
	if api.VoidTransactionHandler == nil {
		api.VoidTransactionHandler = operations.VoidTransactionHandlerFunc(func(params operations.VoidTransactionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.VoidTransaction has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return customMiddleware.NewTraceableHandler(handler)
}
