package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-practice/internal/dto"
	"github.com/go-practice/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user *dto.User
	ctx.BindJSON(&user)

	if user, err := h.userService.CreateUser(ctx, user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		ctx.JSON(http.StatusOK, user)
	}

}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	var users *[]dto.User

	if output, err := h.userService.GetUsers(ctx, users); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		ctx.JSON(http.StatusOK, output)
	}

}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	var user *dto.User
	name := ctx.Param("name")
	if output, err := h.userService.GetUser(ctx, user, name); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		ctx.JSON(http.StatusOK, output)
	}

}
