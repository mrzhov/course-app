package task

import (
	"gorm.io/gorm"
)

type IRepository interface {
	Create(task Task) (Task, error)
	GetList() ([]Task, error)
	GetById(id uint) (Task, error)
	Patch(task Task) (Task, error)
	Delete(id uint) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(task Task) (Task, error) {
	res := r.db.Create(&task)
	if res.Error != nil {
		return Task{}, res.Error
	}
	return task, nil
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

func (r *Repository) Patch(task Task) (Task, error) {
	err := r.db.Save(&task).Error
	return task, err
}

func (r *Repository) Delete(id uint) error {
	err := r.db.Delete(&Task{}, id).Error
	return err
}
