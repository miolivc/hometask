package models

type Task struct {
	ID        int    `json:"id" gorm:"primary_key`
	Name      string `json:"name" gorm:"unique"`
	Level     int    `json:"level" gorm:"not null"`
	Daily     bool   `json:"daily" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"not null" default:"true"`
}
