package forms

type User struct {
	Title     string `json:"title"`
	Author    string `json:"auithor" default:"admin"`
}