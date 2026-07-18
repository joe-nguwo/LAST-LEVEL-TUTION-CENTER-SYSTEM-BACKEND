package models
import(
	"gorm.io/gorm"
)

type Admin struct {
	
	Password  string  `json:"password" gorm:"primaryKey"`
	FName      string `json:"fname"`
	LName      string `json:"lname"`
	Email     string `json:"email" `
	gorm.Model

}