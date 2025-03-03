package gorillax

import (
	"github.com/gorilla/mux"
)

func Middlewares(mdw ...mux.MiddlewareFunc) []mux.MiddlewareFunc {
	return append([]mux.MiddlewareFunc{}, mdw...)
}
