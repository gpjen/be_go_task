package models

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(task Task) (Task, error)
	FindAll() ([]Task, error)
	FindById(id int) (Task, error)
	Update(task Task, id int) (Task, error)
	Delete(task Task) error
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

func (r *repository) Update(task Task) (Task, error) {
	err := r.db.Save(&task).Error
	return task, err
}

func (r *repository) Delete(task Task) error {
	err := r.db.Delete(&task).Error
	return err
}
