package routes

import (
	
	"net/http"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

/*Create Person Struct used for input (not to be confused with Person Struct In
Models, which is used to define a database model)*/
type CreatePerson struct{
	FirstName string `json:"firstName" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Age int `json:"age" binding:"required"`
}



//root endpoint
func GetStart(c *gin.Context){
	c.JSON(200, gin.H{
		"name":"hello",
	})
}

//gets all people in the database
func GetPeople(c *gin.Context) {
	var people []models.Person
	models.DB.Find(&people)

	c.JSON(http.StatusOK, gin.H{"data": people})
}

//adds person to database (test with Insomnia or Postman)
func PostPeople(c *gin.Context) {
	var input CreatePerson
	//checks if the input is valid JSON
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//create a person model object using the fields from CreatePerson
	person := models.Person{FirstName: input.FirstName, LastName: input.LastName, Age: input.Age}
	models.DB.Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": input})
	c.JSON(http.StatusOK, gin.H{"result": person})
	
}

//Deletes a person from the database
func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	var people []models.Person
	
	models.DB.Delete(&people, id)
	c.JSON(http.StatusOK, gin.H{"result":"successfully deleted data"})

}


//Updates a person from the database
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






