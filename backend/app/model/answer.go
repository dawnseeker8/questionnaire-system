package model

import (
	"time"

	"gorm.io/gorm"
)

// Answer represents a user's answer to a question
type Answer struct {
	ID         uint           `json:"id" gorm:"primaryKey;column:ID"`
	ResponseID uint           `json:"response_id" gorm:"not null;column:ResponseID"`
	QuestionID uint           `json:"question_id" gorm:"not null;column:QuestionID"`
	Content    *string        `json:"content" gorm:"type:text;column:Content"` // Text content of the answer
	Score      float64        `json:"score" gorm:"default:0;column:Score"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
