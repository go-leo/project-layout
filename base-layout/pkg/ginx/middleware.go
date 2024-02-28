package ginx

import (
	"github.com/gin-gonic/gin"
)

func Middlewares(mdw ...gin.HandlerFunc) []gin.HandlerFunc {
	return append([]gin.HandlerFunc{}, mdw...)
}
