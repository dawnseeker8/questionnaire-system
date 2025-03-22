package model

import (
	"time"

	"gorm.io/gorm"
)

// QuestionOption represents an option for a question
type QuestionOption struct {
	ID         uint           `json:"id" gorm:"primaryKey;column:ID"`
	QuestionID uint           `json:"question_id" gorm:"not null;column:QuestionID"`
	Text       string         `json:"text" gorm:"size:255;not null;column:Text"`
	Score      float64        `json:"score" gorm:"default:0;column:Score"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
