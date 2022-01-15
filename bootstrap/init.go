package bootstrap

import (
	"context"
	"github.com/google/uuid"
	"github.com/thermeon/boilerplate-golang/configs"
	"github.com/thermeon/boilerplate-golang/infrastructures/server/operations"
	"github.com/thermeon/boilerplate-golang/interfaces/authentication"
	"github.com/thermeon/boilerplate-golang/interfaces/handlers"
	"github.com/thermeon/boilerplate-golang/log"
)

func Init(api *operations.PaymentsServiceAPIAPI) {

	ctx := log.WithTraceableCtx(context.Background(), uuid.New())
	conf := &configs.Config{}
	conf.Init()

	log.Init(&log.Config{
		Level: log.InfoLevel,
	})

	log.Info(ctx, "Start to initialized the payments service.")

	// Attach the custom logger to the server.
	api.Logger = getAPILogger(ctx)

	// Set authenticator to api.
	api.BearerAuthenticator = authentication.NewBearerAuthenticatorWrapper(conf)

	// Set handlers.
	api.AuthorizationTransactionHandler = handlers.GetAuthorizationTransactionHandlerFunc()
}

func getAPILogger(ctx context.Context) func(string, ...interface{}) {
	return func(s string, i ...interface{}) {
		log.Info(ctx, s, i...)
	}
}
