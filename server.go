package main

import (
	"be_go_task/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	// dsn := "root:@tcp(127.0.0.1:3306)/go_task?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// db.AutoMigrate(&models.Task{})

	// task := models.Task{
	// 	Task:     "ui ux",
	// 	Assignor: "syukran reza",
	// 	Dateline: "2022-10-10",
	// }

	// taskRequest := models.NewRepository(db)

	// data, _ := taskRequest.FindById(3)

	// fmt.Println(data)

	router := gin.Default()

	V1 := router.Group("/v1")
	V1.GET("/task", controllers.GetTask)
	V1.GET("/task/:id", controllers.GetTaskById)
	V1.POST("/task", controllers.NewTask)
	V1.PATCH("/task/:id", controllers.UpdateTask)
	V1.DELETE("/task/:id", controllers.DeleteTask)

	router.Run()
}
