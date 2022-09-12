package models

type Person struct{
	ID uint `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
}