package models



type Admin struct {
	Password  string  `json:"password"`
	//Id        string    `json:"id"`
	FName      string `json:"fname"`
	LName      string `json:"lname"`
	Email     string `json:"email"`

}