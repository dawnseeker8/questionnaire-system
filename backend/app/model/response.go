package model

import (
	"time"

	"gorm.io/gorm"
)

// Response represents a user's submission of a form
type Response struct {
	ID         uint           `json:"id" gorm:"primaryKey;column:ID"`
	FormID     uint           `json:"form_id" gorm:"not null;column:FormID"`
	Answers    []Answer       `json:"answers,omitempty" gorm:"foreignKey:ResponseID"`
	SubmittedAt time.Time     `json:"submitted_at" gorm:"column:SubmittedAt"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
