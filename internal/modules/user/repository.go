package user

import (
	"gorm.io/gorm"
)

type IRepository interface {
	Create(user *User) error
	GetList(users *[]User) error
	GetById(user *User, id uint) error
	Patch(user *User) error
	Delete(user *User) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetList(users *[]User) error {
	return r.db.Find(users).Error
}

func (r *Repository) GetById(user *User, id uint) error {
	return r.db.First(user, id).Error
}

func (r *Repository) Patch(user *User) error {
	return r.db.Save(user).Error
}

func (r *Repository) Delete(user *User) error {
	return r.db.Delete(user).Error
}
