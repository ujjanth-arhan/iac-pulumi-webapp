package dto

import "time"

var assignemnt struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"default: null"`
	Points         int       `json:"points" gorm:"default: null;"`
	Num_Of_Attemps int       `json:"num_of_attemps" gorm:"default: null;"`
	Deadline       time.Time `json:"deadline" gorm:"default: null;"`
}
