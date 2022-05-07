package model

type User struct {
	ID    uint64 `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
