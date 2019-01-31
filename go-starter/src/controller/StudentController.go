package controller

import (
	"../service/implementation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudents(context *gin.Context) {
	context.SecureJSON(http.StatusOK, implementation.GetStudents())
}