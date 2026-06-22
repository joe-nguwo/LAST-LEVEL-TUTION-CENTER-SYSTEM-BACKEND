package models

import (
	"gorm.io/gorm"
)

type Books struct{
	gorm.Model
	Name string `json:"name"` 
	Author string `json:"author"`
	Genre string `json:"genre"`
	


}