package entitiy

import "github.com/google/uuid"

type Blog struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	UserID   uuid.UUID `json:"userid" binding:"required"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Likes    int       `json:"likes"`
	Comments []Comment `gorm:"foreignKey:BlogID"`
}
