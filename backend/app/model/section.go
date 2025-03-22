package model

import (
	"time"

	"gorm.io/gorm"
)

// Section represents a group of related questions in a form
type Section struct {
	ID          uint           `json:"id" gorm:"primaryKey;column:ID"`
	FormID      uint           `json:"form_id" gorm:"not null;column:FormID"`
	Title       string         `json:"title" gorm:"size:255;not null;column:Title"`
	Description string         `json:"description" gorm:"type:text;column:Description"`
	OrderIndex  int            `json:"order_index" gorm:"default:0;column:OrderIndex"`
	TotalScore  float64        `json:"total_score" gorm:"default:0;column:TotalScore"`
	Questions   []Question     `json:"questions,omitempty" gorm:"foreignKey:SectionID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
