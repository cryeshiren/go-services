package router

import (
	e "../../error"
	h "../router/http"
	"github.com/gin-gonic/gin"
)

type RouterMetaData struct {
	method h.Method
	controller func(c *gin.Context)
}

var routerMap = make(map[string]RouterMetaData)

func RegistRouter(router *gin.Engine) {
	for url, metaData := range routerMap {
		switch metaData.method {
			case h.Get:
				router.GET(url, metaData.controller)
			case h.Post:
			router.POST(url, metaData.controller)
			case h.Put:
				router.PUT(url, metaData.controller)
			case h.Delete:
				router.DELETE(url, metaData.controller)
			case h.Patch:
				router.PATCH(url, metaData.controller)
		}
	}
}

func Router(method h.Method, url string, action func(c *gin.Context)) error{
	if _, ok := routerMap[url]; ok {
		return e.RouterExistError{
			Url: url,
		}
	}
	routerMap[url] = RouterMetaData{
		method:method,
		controller:action,
	}
	return nil
}