package server

import (
	"net/http"

	"github.com/go-chi/chi"
	cmiddleware "github.com/go-chi/chi/middleware"
	"github.com/kitabisa/gohello/internal/app/commons"
	"github.com/kitabisa/gohello/internal/app/handler"
	"github.com/kitabisa/gohello/version"
	phttp "github.com/kitabisa/perkakas/v2/http"
	pmiddleware "github.com/kitabisa/perkakas/v2/middleware"
	pstructs "github.com/kitabisa/perkakas/v2/structs"
)

// Router a chi mux
func Router(opt handler.HandlerOption) *chi.Mux {
	handlerCtx := phttp.NewContextHandler(pstructs.Meta{
		Version: version.Version,
		Status:  "stable", //TODO: ask infra if this is used
		APIEnv:  version.Environment,
	})
	commons.InjectErrors(&handlerCtx)

	logMiddleware := pmiddleware.NewHttpRequestLogger(opt.Logger)

	r := chi.NewRouter()
	r.Use(logMiddleware)
	r.Use(cmiddleware.Recoverer)

	// the handler
	phandler := phttp.NewHttpHandler(handlerCtx)

	healthCheckHandler := handler.HealthCheckHandler{}

	healthCheckHandler.HandlerOption = opt
	healthCheckHandler.Handler = phandler(healthCheckHandler.HealthCheck)

	helloHandler := handler.HelloHandler{}
	helloHandler.HandlerOption = opt
	helloHandler.Handler = phandler(helloHandler.SayHello)

	// Setup your routing here
	r.Method(http.MethodGet, "/health_check", healthCheckHandler)
	r.Method(http.MethodGet, "/hello", helloHandler)
	return r
}

// TODO: func authRouter()
