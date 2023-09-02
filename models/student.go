package models

type Student struct {
	Rev string `json:"_rev"`
	Name string `json:"name"`
	Age int `json:"age"`
	ID string `json:"id"`
	Classroom string `json:"classroom"`
}