package model

import (
	"time"

	"gorm.io/gorm"
)

// Form represents a questionnaire or survey form
type Form struct {
	ID          uint           `json:"id" gorm:"primaryKey;column:ID"`
	Title       string         `json:"title" gorm:"size:255;not null;column:Title"`
	Description string         `json:"description" gorm:"type:text;column:Description"`
	Sections    []Section      `json:"sections,omitempty" gorm:"foreignKey:FormID"`
	Responses   []Response     `json:"responses,omitempty" gorm:"foreignKey:FormID"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:CreatedAt"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"column:UpdatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
