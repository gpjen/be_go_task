package models

import "time"

type Task struct {
	Id        uint   `gorm:"primary_key;autoIncrement;not null"`
	Task      string `gorm:"size:255"`
	Assignor  string `gorm:"size:100"`
	Dateline  string `gorm:"size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
