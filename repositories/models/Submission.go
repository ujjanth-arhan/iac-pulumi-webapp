package models

import "time"

type Submission struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	AccountId      string    `json:"account_id" gorm:"default: null; not null"`
	AssignmentId   string    `json:"assignment_id" gorm:"default: null; not null"`
	SubmissionUrl  string    `json:"submission_url" gorm:"default: null; not null"`
	SubmissionDate time.Time `json:"account_created" gorm:"autoCreateTime;"`
}
