package task

import (
	"gorm.io/gorm"
)

type IRepository interface {
	Create(task *Task) error
	GetList(tasks *[]Task) error
	GetById(task *Task, id uint) error
	Patch(task *Task) error
	Delete(task *Task) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(task *Task) error {
	return r.db.Create(task).Error
}

func (r *Repository) GetList(tasks *[]Task) error {
	return r.db.Find(tasks).Error
}

func (r *Repository) GetById(task *Task, id uint) error {
	return r.db.First(task, id).Error
}

func (r *Repository) Patch(task *Task) error {
	return r.db.Save(task).Error
}

func (r *Repository) Delete(task *Task) error {
	return r.db.Delete(task).Error
}
