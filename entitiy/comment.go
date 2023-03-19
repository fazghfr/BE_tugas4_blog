package entitiy

import "github.com/google/uuid"

type Comment struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	BlogID  uuid.UUID `json:"blogid" binding:"required"`
	Content string    `json:"content"`
}
