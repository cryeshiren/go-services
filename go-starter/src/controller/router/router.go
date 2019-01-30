package router

import (
	h "../router/http"
	"github.com/gin-gonic/gin"
)

type routerMetaData struct {
	method h.Method
	controller func(c *gin.Context)
}

var routerMap map[string]func(c *gin.Context)

func RegistRouter(router *gin.Engine) {
}