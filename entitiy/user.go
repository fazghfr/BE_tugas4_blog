package entitiy

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Password string    `json:"password"`
}
