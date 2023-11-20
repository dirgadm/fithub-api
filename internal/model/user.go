package model

import (
	"time"
)

type User struct {
	Id        int `gorm:"primaryKey;autoIncrement:true"`
	Email     string
	Password  string
	Name      string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) TableName() string {
	return "user"
}
