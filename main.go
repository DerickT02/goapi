package main

import (
	"fmt"
 	"github.com/gin-gonic/gin"
	"restapi/routes"
	"restapi/models"
)


func main() {
	//Declaring Routes and running server
	fmt.Println("Hello World")
	r := gin.Default()

	//connect to Database using connection function from models package
	models.DBConn()

	r.GET("/", routes.GetStart)
	r.GET("/people", routes.GetPeople)
	r.PATCH("/people/:id", routes.UpdatePerson)
	r.POST("/people", routes.PostPeople)
	r.DELETE("/people/:id", routes.DeletePerson)

	r.Run("localhost:3000")
}