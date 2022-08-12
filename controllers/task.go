package controllers

import (
	"be_go_task/config"
	"be_go_task/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// create
func NewTask(c *gin.Context) {

	var dataRequest TaskRequest
	err := c.ShouldBindJSON(&dataRequest)
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

	db, _ := config.DbConn()
	TaskRequest := models.NewRepository(db)

	data := models.Task{
		Task:     dataRequest.Task,
		Assignor: dataRequest.Assignor,
		Dateline: dataRequest.Dateline,
	}

	TaskRequest.Create(data)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "create new task",
	})
}

// read all
func GetTask(c *gin.Context) {

	db, err := config.DbConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}

	taskRequest := models.NewRepository(db)
	getData, err := taskRequest.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}

	var data []TaskResponse

	for _, t := range getData {
		d := TaskResponse{
			Id:       t.Id,
			Task:     t.Task,
			Assignor: t.Assignor,
			Dateline: t.Dateline,
		}
		data = append(data, d)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "get task",
		"data":    data,
	})
}

// read by id
func GetTaskById(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	db, err := config.DbConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}

	taskRequest := models.NewRepository(db)
	getData, err := taskRequest.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	} else if getData.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprint("no data id ", id),
		})
		return
	}

	data := TaskResponse{
		Id:       getData.Id,
		Task:     getData.Task,
		Assignor: getData.Assignor,
		Dateline: getData.Dateline,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprint("get task by id ", id),
		"data":    data,
	})
}

// update task
func UpdateTask(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	var dataRequest TaskRequest

	err := c.ShouldBindJSON(&dataRequest)
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

	db, err := config.DbConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}
	taskRequest := models.NewRepository(db)
	getData, _ := taskRequest.FindById(id)

	if getData.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprint("no data id ", id),
		})
		return
	}

	getData.Task = dataRequest.Task
	getData.Assignor = dataRequest.Assignor
	getData.Dateline = dataRequest.Dateline

	taskRequest.Update(getData)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprint("update task id ", id),
	})
}

// delete task
func DeleteTask(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	db, err := config.DbConn()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	TaskRequest := models.NewRepository(db)
	data, _ := TaskRequest.FindById(id)

	if data.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprint("no data id ", id),
		})
		return
	}

	TaskRequest.Delete(data)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": fmt.Sprint("delete task id ", id),
	})
}
