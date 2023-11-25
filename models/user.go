package models

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
