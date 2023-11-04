package model

type User struct {
	ID       uint   `gorm:"primarykey"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Tales    []Tale // One-to-Many relationship: a User can have multiple Tales
}

type UserUpdate struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword	"`
}
