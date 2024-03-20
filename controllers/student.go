package controllers

import (
	"restapi/initializers"
	"restapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var Body struct {
		Name string
		Age int
	}
	c.Bind(&Body)
	if Body.Name == "" || Body.Age==0 {
		c.Status(400)
		return;
	}
	student := models.Student{Name: Body.Name, Age: Body.Age}
	result := initializers.DB.Create(&student)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(201, gin.H{"student": student})
}

func GetAll(c *gin.Context) {
	var students []models.Student
	initializers.DB.Find(&students)
	c.JSON(200, gin.H{"students": students})
}

func GetOne(c *gin.Context) {
	id := c.Param("id")
	studentId, err := strconv.Atoi(id)
	if err != nil {
		c.Status(400)
		return
	}
	var student models.Student
	initializers.DB.First(&student, studentId)
	c.JSON(200, gin.H{"student": student})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	studentId, err := strconv.Atoi(id)
	if err != nil {
		c.Status(400)
		return
	}
	var student models.Student
	initializers.DB.First(&student, studentId)

	var Body struct {
		Name string
		Age int
	}

	c.Bind(&Body)

	student.Name = Body.Name
	student.Age = Body.Age

	initializers.DB.Save(&student)

	c.JSON(200, gin.H{"student": student})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	studentId, err := strconv.Atoi(id)
	if err != nil {
		c.Status(400)
		return
	}
	student := models.Student{}
	student.ID = uint(studentId)
	initializers.DB.Delete(&student)
	c.JSON(200, gin.H{"message": "deleted student with id " + id})
}