package controller

import (
	r "../controller/router"
	h "../controller/router/http"
	"../service/implementation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RR() {
	_ = r.Router(h.Get, "/students", getStudents)
}

func getStudents(context *gin.Context) {
	context.SecureJSON(http.StatusOK, implementation.GetStudents())
}