package service

import (
	"context"

	"github.com/go-practice/internal/dto"
	"github.com/go-practice/internal/repository"
	"github.com/opentracing/opentracing-go"
)

type UserService struct {
	repository repository.IRepository
}

func NewUserService(repository repository.IRepository) *UserService {
	return &UserService{repository: repository}
}

// create user
func (s *UserService) CreateUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserService.CreateUser")
	defer span.Finish()

	user, err := s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// get users
func (s *UserService) GetUsers(ctx context.Context, users *[]dto.User) (*[]dto.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserService.GetUsers")
	defer span.Finish()

	output, err := s.repository.GetUsers(ctx, users)
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return nil, err
	}
	// c.JSON(http.StatusOK, user)

	return output, nil
}

// get user by id
func (s *UserService) GetUser(ctx context.Context, user *dto.User, name string) (*dto.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserService.GetUser")
	defer span.Finish()

	output, err := s.repository.GetUser(ctx, user, name)
	if err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	c.AbortWithStatus(http.StatusNotFound)
		return nil, err
	}

	// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	// return
	// c.JSON(http.StatusOK, user)

	return output, nil
}

// // update user
// func (s *UserService) UpdateUser(c *gin.Context, user *dto.User) {
// 	var user models.User
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	err := models.GetUser(repository.Db, &user, id)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			c.AbortWithStatus(http.StatusNotFound)
// 			return
// 		}

// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.BindJSON(&user)
// 	err = models.UpdateUser(repository.Db, &user)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// // delete user
// func (s *UserService) DeleteUser(c *gin.Context) {
// 	var user models.User
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	err := models.DeleteUser(repository.Db, &user, id)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }
