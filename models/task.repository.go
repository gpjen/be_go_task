package models

import "gorm.io/gorm"

type Repository interface {
	Create(task Task) (Task, error)
	FindAll() ([]Task, error)
	FindById(id int) (Task, error)
	Update(task Task, id int) (Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *repository) FindById(id int) (Task, error) {
	var task Task
	err := r.db.Find(&task, id).Error
	return task, err
}

func (r *repository) Create(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

// func (r *repository) Update(newTask Task, id int) (Task, error) {
// 	var task Task

// 	err := r.db.Model(&task).Update()

// }
