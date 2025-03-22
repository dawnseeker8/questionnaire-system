package model

import (
	"time"

	"gorm.io/gorm"
)

// Question represents a question in a section
type Question struct {
	ID         uint             `json:"id" gorm:"primaryKey;column:ID"`
	SectionID  uint             `json:"section_id" gorm:"not null;column:SectionID"`
	Type       string           `json:"type" gorm:"size:50;not null;column:Type"` // radio, checkbox, text
	Text       string           `json:"text" gorm:"type:text;not null;column:Text"`
	IsRequired bool             `json:"is_required" gorm:"default:false;column:IsRequired"`
	OrderIndex int              `json:"order_index" gorm:"default:0;column:OrderIndex"`
	Options    []QuestionOption `json:"options,omitempty" gorm:"foreignKey:QuestionID"`
	Answers    []Answer         `json:"-" gorm:"foreignKey:QuestionID"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	DeletedAt  gorm.DeletedAt   `json:"-" gorm:"index"`
}
