package controller

import (
	"../service/implementation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//TODO: refactor error

// GetStudentById godoc
// @Tags Student
// @Summary Show student info
// @Description get students by id
// @ID student id
// @Accept  json
// @Produce  json
// @Param id path int true "Student ID"
// @Success 200 {Student} Student
// @Failure 400 {interface{}} string "Bad Request"
// @Failure 404 {interface{}} string "Not Found"
// @Failure 500 {interface{}} string "Server Error"
// @Router /students/{id} [get]
func GetStudentById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	context.SecureJSON(http.StatusOK, implementation.GetStudentById(id))
}