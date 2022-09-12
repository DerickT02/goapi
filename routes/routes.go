package routes

import "restapi/models"
import "github.com/gin-gonic/gin"
import "net/http"



type CreatePerson struct{
	FirstName string `json:"firstName" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Age int `json:"age" binding:"required"`
}



func GetStart(c *gin.Context){
	c.JSON(200, gin.H{
		"name":"hello",
	})
}


func GetPeople(c *gin.Context) {
	var people []models.Person
	models.DB.Find(&people)

	c.JSON(http.StatusOK, gin.H{"data": people})
}

func PostPeople(c *gin.Context) {
	var input CreatePerson

	//checks if the input is valid JSON
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person := models.Person{FirstName: input.FirstName, LastName: input.LastName, Age: input.Age}
	models.DB.Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": person})
	
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	var people []models.Person
	
	models.DB.Delete(&people, id)
	c.JSON(http.StatusOK, gin.H{"result":"successfully deleted data"})

}

func UpdatePerson(c *gin.Context) {
	var input CreatePerson
	var person models.Person
	
	//find the model and then update it based on the input
	    

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



 


	if err := models.DB.Where("id = ?",  c.Param("id")).First(&person).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    	return
	}



	models.DB.Model(&person).Updates(input)
	
	
	


	c.JSON(http.StatusOK, gin.H{"result": person})


}






