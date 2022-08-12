package controllers

type TaskRequest struct {
	Task     string `json:"task" binding:"required"`
	Assignor string `json:"assignor" binding:"required"`
	Dateline string `json:"dateline" binding:"required"`
}

type TaskResponse struct {
	Id       uint   `json:"id"`
	Task     string `json:"task"`
	Assignor string `json:"assignor"`
	Dateline string `json:"dateline"`
}
