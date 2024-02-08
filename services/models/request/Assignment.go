package request

import "time"

type Assignment struct {
	Name           string    `gorm:"default: null; not null"`
	Points         int       `gorm:"default: null; not null"`
	Num_Of_Attemps int       `gorm:"default: null; not null"`
	Deadline       time.Time `gorm:"default: null; not null"`
}
