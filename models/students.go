package models

import(
	"gorm.io/gorm"
)


type Students struct{
	
	FirstName string `json:"fname"`
	MiddleName string `json:"m-name"`
	LastName string `json:"lname"`
	Email string `json:"email"  gorm:"unique"`
	Contact string `json:"contact" gorm:"unique"`
	Address string `json:"address"`
	gorm.Model

}