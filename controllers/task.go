package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TaskRequest struct {
	Task     string `json:"task" binding:"required"`
	Assignor string `json:"assignor" binding:"required"`
	Dateline string `json:"dateline" binding:"required"`
}

// create
func NewTask(c *gin.Context) {

	var data TaskRequest
	err := c.ShouldBindJSON(&data)
	if err != nil {
		var errMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("error on field %s, condition %s", e.Field(), e.ActualTag())
			errMessages = append(errMessages, errMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": errMessages,
		})
		return

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
		var errMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("error on field %s, condition %s", e.Field(), e.ActualTag())
			errMessages = append(errMessages, errMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": errMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprint("update task id ", id),
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
