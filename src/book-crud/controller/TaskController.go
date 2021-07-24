package controller

import (
	"book-crud/constant"
	"book-crud/model"
	"book-crud/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateTaskInput struct
type CreateTaskInput struct {
	AssignedTo  string `json:"assignedTo"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

// UpdateTaskInput struct
type UpdateTaskInput struct {
	AssignedTo  string `json:"assignedTo"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

// FindTasks method:GET endpoint:/tasks, desc:Get all tasks
func FindTasks(c *gin.Context) {
	var tasks []model.Task

	db := c.MustGet("db").(*gorm.DB)
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{
		"count": len(tasks),
		"data":  tasks,
	})
}

// CreateTask method:POST, endpoint:/tasks, desc:Create new task
func CreateTask(c *gin.Context) {
	//Validate input
	var input CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deadline, err := time.Parse(constant.TIME_PATTERN, util.ReformatStringDate(input.Deadline))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	task := model.Task{AssignedTo: input.AssignedTo, Description: input.Description, Deadline: deadline}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// FindTask method:GET, endpoint:/tasks/:id, desc:Find a task
func FindTask(c *gin.Context) {
	var task model.Task
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// UpdateTask method:PATCH endpoint:/tasks/:id, desc:Update a task
func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var task model.Task

	if err := db.Where("id=?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input of Task
	var input UpdateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deadline, err := time.Parse(constant.TIME_PATTERN, util.ReformatStringDate(input.Deadline))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput model.Task
	updatedInput.AssignedTo = input.AssignedTo
	updatedInput.Description = input.Description
	updatedInput.Deadline = deadline

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DeleteTask method:DELETE,  endpoint:/tasks/:id, desc:Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book model.Task

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
