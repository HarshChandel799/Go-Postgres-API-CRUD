package main

import (
	"go_postgres/controllers"
	"go_postgres/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVvariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/student_add", controllers.StudentCreate)
	r.GET("/student_get", controllers.StudentsGet)
	r.GET("/student_get/:id", controllers.StudentGetID)
	r.PUT("/student_update/:id", controllers.StudentUpdate)
	r.DELETE("/student_delete/:id", controllers.StudentDelete)
	r.Run()
}
