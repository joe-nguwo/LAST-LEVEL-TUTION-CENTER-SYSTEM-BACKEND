package models

import (
	"gorm.io/gorm"
)

type Books struct{
	
	Name string `json:"name"` 
	Author string `json:"author"`
	Genre string `json:"genre"`
	gorm.Model
	


}