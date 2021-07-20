package controller

import (
	"BookCrud/constant"
	"BookCrud/model"
	"BookCrud/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssignedTo  string `json:"assignedTo"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

type UpdateTaskInput struct {
	AssignedTo  string `json:"assignedTo"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []model.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{
		"count": len(tasks),
		"data":  tasks,
	})
}

// POST /tasks
// Create new task
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

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task model.Task
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var task model.Task

	if err := db.Where("id=?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
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
