package controllers

import (
	"go_postgres/initializers"
	"go_postgres/models"

	"github.com/gin-gonic/gin"
)

func StudentCreate(c *gin.Context) {

	var body struct {
		Name    string
		Age     string
		Class   string
		Section string
	}
	c.Bind(&body)
	student := models.Student{Name: body.Name, Age: body.Age, Class: body.Class, Section: body.Section}
	result := initializers.DB.Create(&student)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"student": student,
	})
}

func StudentsGet(c *gin.Context) {
	var students []models.Student
	result := initializers.DB.Find(&students)

	if result.Error != nil {
		c.JSON(200, gin.H{
			"message": "No Students Found",
		})
	}
	c.JSON(200, gin.H{
		"students": students,
	})
}

func StudentGetID(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	result := initializers.DB.Find(&student, id)

	if result.Error != nil {
		c.JSON(200, gin.H{
			"message": "No Student Found",
		})
	}
	c.JSON(200, gin.H{
		"student": student,
	})
}

func StudentUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Name    string
		Age     string
		Class   string
		Section string
	}
	c.Bind(&body)
	var student models.Student
	initializers.DB.Find(&student, id)
	initializers.DB.Model(&student).Updates(models.Student{
		Name:    body.Name,
		Age:     body.Age,
		Class:   body.Class,
		Section: body.Section,
	})
	c.JSON(200, gin.H{
		"student": student,
	})
}

func StudentDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Student{}, id)
	c.Status(200)
}
