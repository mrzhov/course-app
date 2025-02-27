package task

import (
	"gorm.io/gorm"
)

type IRepository interface {
	Create(task *Task) error
	GetList() ([]Task, error)
	GetById(id uint) (Task, error)
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

func (r *Repository) GetList() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *Repository) GetById(id uint) (Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	return task, err
}

func (r *Repository) Patch(task *Task) error {
	return r.db.Save(task).Error
}

func (r *Repository) Delete(task *Task) error {
	return r.db.Delete(task).Error
}
