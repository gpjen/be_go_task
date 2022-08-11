package models

import "time"

type Task struct {
	Id        int
	Task      string
	Assignor  string
	Dateline  string
	CreatedAt time.Time
	UpdateAt  time.Time
}
