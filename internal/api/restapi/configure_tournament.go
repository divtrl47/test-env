// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"kafka-playground/internal/api/restapi/operations"
	"kafka-playground/internal/api/restapi/operations/auth"
	"kafka-playground/internal/api/restapi/operations/hints"
	"kafka-playground/internal/api/restapi/operations/tasks"
	"kafka-playground/internal/api/restapi/operations/tournaments"
)

//go:generate swagger generate server --target ../../api --name Tournament --spec ../../../docs/api/openapi.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.TournamentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TournamentAPI) http.Handler {
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

	if api.TournamentsGetTournamentsHandler == nil {
		api.TournamentsGetTournamentsHandler = tournaments.GetTournamentsHandlerFunc(func(params tournaments.GetTournamentsParams) middleware.Responder {
			return middleware.NotImplemented("operation tournaments.GetTournaments has not yet been implemented")
		})
	}
	if api.TournamentsGetTournamentsTournamentIDHandler == nil {
		api.TournamentsGetTournamentsTournamentIDHandler = tournaments.GetTournamentsTournamentIDHandlerFunc(func(params tournaments.GetTournamentsTournamentIDParams) middleware.Responder {
			return middleware.NotImplemented("operation tournaments.GetTournamentsTournamentID has not yet been implemented")
		})
	}
	if api.TasksGetTournamentsTournamentIDTasksTaskIDHandler == nil {
		api.TasksGetTournamentsTournamentIDTasksTaskIDHandler = tasks.GetTournamentsTournamentIDTasksTaskIDHandlerFunc(func(params tasks.GetTournamentsTournamentIDTasksTaskIDParams) middleware.Responder {
			return middleware.NotImplemented("operation tasks.GetTournamentsTournamentIDTasksTaskID has not yet been implemented")
		})
	}
	if api.HintsGetTournamentsTournamentIDTasksTaskIDHintsHintIDHandler == nil {
		api.HintsGetTournamentsTournamentIDTasksTaskIDHintsHintIDHandler = hints.GetTournamentsTournamentIDTasksTaskIDHintsHintIDHandlerFunc(func(params hints.GetTournamentsTournamentIDTasksTaskIDHintsHintIDParams) middleware.Responder {
			return middleware.NotImplemented("operation hints.GetTournamentsTournamentIDTasksTaskIDHintsHintID has not yet been implemented")
		})
	}
	if api.TasksPatchTournamentsTournamentIDTasksTaskIDHandler == nil {
		api.TasksPatchTournamentsTournamentIDTasksTaskIDHandler = tasks.PatchTournamentsTournamentIDTasksTaskIDHandlerFunc(func(params tasks.PatchTournamentsTournamentIDTasksTaskIDParams) middleware.Responder {
			return middleware.NotImplemented("operation tasks.PatchTournamentsTournamentIDTasksTaskID has not yet been implemented")
		})
	}
	if api.HintsPatchTournamentsTournamentIDTasksTaskIDHintsHintIDHandler == nil {
		api.HintsPatchTournamentsTournamentIDTasksTaskIDHintsHintIDHandler = hints.PatchTournamentsTournamentIDTasksTaskIDHintsHintIDHandlerFunc(func(params hints.PatchTournamentsTournamentIDTasksTaskIDHintsHintIDParams) middleware.Responder {
			return middleware.NotImplemented("operation hints.PatchTournamentsTournamentIDTasksTaskIDHintsHintID has not yet been implemented")
		})
	}
	if api.AuthPostLoginHandler == nil {
		api.AuthPostLoginHandler = auth.PostLoginHandlerFunc(func(params auth.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostLogin has not yet been implemented")
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
