// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Konstantinov-Innokentii/mrmixr/api/restapi/operations"
)

//go:generate swagger generate server --target ../../api --name Mrmixr --spec ../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.MrmixrAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MrmixrAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CreateCheckHandler == nil {
		api.CreateCheckHandler = operations.CreateCheckHandlerFunc(func(params operations.CreateCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateCheck has not yet been implemented")
		})
	}
	if api.DeleteCheckHandler == nil {
		api.DeleteCheckHandler = operations.DeleteCheckHandlerFunc(func(params operations.DeleteCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteCheck has not yet been implemented")
		})
	}
	if api.ListChecksHandler == nil {
		api.ListChecksHandler = operations.ListChecksHandlerFunc(func(params operations.ListChecksParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListChecks has not yet been implemented")
		})
	}
	if api.ListGithubRepositoriesHandler == nil {
		api.ListGithubRepositoriesHandler = operations.ListGithubRepositoriesHandlerFunc(func(params operations.ListGithubRepositoriesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListGithubRepositories has not yet been implemented")
		})
	}
	if api.RetrieveCheckHandler == nil {
		api.RetrieveCheckHandler = operations.RetrieveCheckHandlerFunc(func(params operations.RetrieveCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.RetrieveCheck has not yet been implemented")
		})
	}
	if api.UpdateCheckHandler == nil {
		api.UpdateCheckHandler = operations.UpdateCheckHandlerFunc(func(params operations.UpdateCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateCheck has not yet been implemented")
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
	return handler
}
