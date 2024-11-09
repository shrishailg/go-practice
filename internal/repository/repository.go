package repository

import (
	"context"

	"github.com/go-practice/internal/dto"
	"gorm.io/gorm"
)

type IRepository interface {
	Create(ctx context.Context, user *dto.User) (*dto.User, error)
	Update(ctx context.Context, user *dto.User) (*dto.User, error)
	GetUsers(ctx context.Context, users *[]dto.User) (*[]dto.User, error)
	GetUser(ctx context.Context, user *dto.User, name string) (*dto.User, error)
	Delete(ctx context.Context, user *dto.User, id int) (*dto.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// create a user
func (r *Repository) Create(ctx context.Context, user *dto.User) (*dto.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// update user
func (r *Repository) Update(ctx context.Context, user *dto.User) (*dto.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUsers(ctx context.Context, users *[]dto.User) (*[]dto.User, error) {
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// get user by id
func (r *Repository) GetUser(ctx context.Context, user *dto.User, name string) (*dto.User, error) {
	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// delete user
func (r *Repository) Delete(ctx context.Context, user *dto.User, id int) (*dto.User, error) {
	err := r.db.Where("id = ?", id).Find(&user).Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
