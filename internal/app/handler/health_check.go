package handler

import (
	"net/http"
)

// HealthCheckHandler object for health check handler
type HealthCheckHandler struct {
	HandlerOption
	http.Handler
}

// HealthCheck checking if all work well
func (h HealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) (data interface{}, pageToken *string, err error) {
	return
}
