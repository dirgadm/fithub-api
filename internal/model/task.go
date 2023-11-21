package model

import (
	"time"
)

type Task struct {
	Id          int `gorm:"primaryKey;autoIncrement:true"`
	Title       string
	Description string
	Status      string
	UserId      int
	Duedate     time.Time
}

func (m *Task) TableName() string {
	return "task"
}
