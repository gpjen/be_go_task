package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskRequest struct {
	Task     string
	Assignor string
	Dateline string
}

// create
func NewTask(c *gin.Context) {
	var data TaskRequest

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "create new task",
		"data":    data,
	})
}

// read all
func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "get task",
	})
}

// read by id
func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprint("get task by id ", id),
	})
}

// update task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var data TaskRequest

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprint("Update task id ", id),
		"data":    data,
	})
}

// delete task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprint("delete task id ", id),
	})
}
