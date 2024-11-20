package repository

import (
	"context"

	"github.com/go-practice/internal/dto"
	"github.com/go-practice/internal/models"
	"gorm.io/gorm"
)

type Repository[M models.MODEL] interface {
	Create(ctx context.Context, model M) (M, error)
	Update(ctx context.Context, model M) (M, error)
	GetUsers(ctx context.Context, model []M) ([]M, error)
	GetUser(ctx context.Context, model M, name string) (M, error)
	Delete(ctx context.Context, model M, id int) (M, error)
}

type RepositoryImpl[M models.MODEL] struct {
	db *gorm.DB
}

func NewRepository[M models.MODEL](db *gorm.DB) *RepositoryImpl[M] {
	return &RepositoryImpl[M]{db: db}
}

// create a user
func (r *RepositoryImpl[M]) Create(ctx context.Context, model M) (M, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// update user
func (r *RepositoryImpl[M]) Update(ctx context.Context, model M) (*dto.User, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *RepositoryImpl[M]) GetUsers(ctx context.Context, model []M) ([]M, error) {
	err := r.db.Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// get user by id
func (r *RepositoryImpl[M]) GetUser(ctx context.Context, model M, name string) (M, error) {
	err := r.db.Where("name = ?", name).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

// delete user
func (r *RepositoryImpl[M]) Delete(ctx context.Context, model M, id int) (M, error) {
	err := r.db.Where("id = ?", id).Find(&model).Delete(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}
