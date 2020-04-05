package handler

import "net/http"

type HelloHandler struct {
	HandlerOption
	http.Handler
}

func (h HelloHandler) SayHello(w http.ResponseWriter, r *http.Request) (data interface{}, pageToken *string, err error) {
	data = h.Services.Hello.SayHelloFromConfig(r.Context())

	return
}
