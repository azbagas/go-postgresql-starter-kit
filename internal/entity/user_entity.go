package entity

import "time"

// User is a struct that represents a user entity
type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Token     string    `gorm:"column:token"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Contacts  []Contact `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
