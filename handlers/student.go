package handlers

import (
	"DB_students/models"

	"net/http"

	"github.com/labstack/echo/v4"

	"DB_students/utilities"
)

func GetAllStudents(c echo.Context) error {
	var students []models.Student

	db.Preload("Direction").Find(&students)

	return c.JSON(http.StatusOK, students)
}

func GetStudentId(c echo.Context) error {
	var student models.Student

	db.Preload("Direction").Take(&student, c.Param("id"))

	if student.ID == 0 {
		return c.String(http.StatusNotFound, "Not found student")
	}

	return c.JSON(http.StatusOK, student)
}

func CreateStudent(c echo.Context) error {
	student := models.Student{
		FullName: c.FormValue("fullName"),
		Age:      utilities.StringToUint(c.FormValue("age")),
	}

	db.Create(&student)

	return c.String(http.StatusCreated, "Student create")
}

func UpdateStudent(c echo.Context) error {
	var student models.Student

	db.Take(&student, c.Param("id"))

	if student.ID == 0 {
		return c.String(http.StatusNotFound, "Not found student")
	}

	student.FullName = c.FormValue("fullName")
	student.Age = utilities.StringToUint(c.FormValue("age"))

	db.Save(&student)

	return c.JSON(http.StatusOK, student)
}

func DeleteStudent(c echo.Context) error {
	var student models.Student

	db.Take(&student, c.Param("id"))

	if student.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found student")
	}

	db.Delete(&student)

	return c.String(http.StatusOK, "Delete student")
}
