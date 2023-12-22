package models

import (
	"github.com/google/uuid"
	"time"
)

type Comm struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	User      uuid.UUID `gorm:"not null" json:"user,omitempty"`
	Post      uuid.UUID `gorm:"not null" json:"post,omitempty"`
	Content   string    `gorm:"not null" json:"content,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
}
type CreateCommRequest struct {
	User      string    `json:"user,omitempty"`
	Post      string    `json:"post,omitempty"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
