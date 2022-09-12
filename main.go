package main

import (
	"fmt"
 	"github.com/gin-gonic/gin"
	"restapi/routes"
	"restapi/models"
)


func main() {
	fmt.Println("Hello World")
	r := gin.Default()

	models.DBConn()

	r.GET("/", routes.GetStart)
	r.GET("/people", routes.GetPeople)
	r.GET("/people/:id", routes.UpdatePerson)
	r.POST("/people", routes.PostPeople)
	r.DELETE("/people/:id", routes.DeletePerson)

	r.Run("localhost:3000")
}