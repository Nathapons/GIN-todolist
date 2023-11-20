package forms

type Blog struct {
	Title     string `json:"title"`
	Author    string `json:"author" default:"admin"`
}