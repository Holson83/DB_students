package main

import (
	"DB_students/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/students", handlers.GetAllStudents)
	e.GET("/students/:id", handlers.GetStudentId)
	e.POST("/students", handlers.CreateStudent)
	e.PUT("/students/:id", handlers.UpdateStudent)
	e.DELETE("/students/:id", handlers.DeleteStudent)

	e.Logger.Fatal(e.Start(":8080"))
}
