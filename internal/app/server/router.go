package server

import (
	"net/http"

	"github.com/go-chi/chi"
	cmiddleware "github.com/go-chi/chi/middleware"
	phttp "github.com/kitabisa/perkakas/v2/http"
	pmiddleware "github.com/kitabisa/perkakas/v2/middleware"
	pstructs "github.com/kitabisa/perkakas/v2/structs"
	"github.com/zackijack/gohello/internal/app/commons"
	"github.com/zackijack/gohello/internal/app/handler"
	"github.com/zackijack/gohello/version"
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
