package models

import "time"

type Assignment struct {
	ID                 string    `json:"id" gorm:"primaryKey"`
	Name               string    `json:"name" gorm:"default: null; not null"`
	Points             int       `json:"points" gorm:"default: null; not null"`
	Num_Of_Attemps     int       `json:"num_of_attemps" gorm:"default: null; not null"`
	Deadline           time.Time `json:"deadline" gorm:"default: null; not null"`
	Assignment_Created time.Time `json:"assignment_created" gorm:"autoCreateTime"`
	Assignment_Updated time.Time `json:"assignment_updated" gorm:"autoUpdateTime"`
}
