package models

type Todo struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}
