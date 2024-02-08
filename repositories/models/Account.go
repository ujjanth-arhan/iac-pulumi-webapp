package models

import "time"

type Account struct {
	ID              string       `json:"id" gorm:"primaryKey"`
	First_Name      string       `json:"first_name" gorm:"default: null; not null"`
	Last_Name       string       `json:"last_name" gorm:"default: null; not null"`
	Email           string       `json:"email" gorm:"default: null; not null; unique"`
	Password        string       `json:"password" gorm:"default: null; not null"`
	Account_Created time.Time    `json:"account_created" gorm:"autoCreateTime;"`
	Account_Updated time.Time    `json:"account_updated" gorm:"autoUpdateTime;"`
	Assignments     []Assignment `json:"assignments" gorm:"many2many:account_assignments"`
}
