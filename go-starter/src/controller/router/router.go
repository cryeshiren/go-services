package router

import (
	"../../controller"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(engine *gin.Engine){
	students := engine.Group("/students")
	{
		students.GET("/:id", controller.GetStudents)
	}
}