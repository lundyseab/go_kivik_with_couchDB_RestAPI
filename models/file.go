package models

type FileDoc struct {
	ID string `json:"_id"`
	REV string `json:"_rev"`
	Data string `json:"data"`
	Name string `json:"name"`
}