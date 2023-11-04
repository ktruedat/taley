package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/ktruedat/taleBack/internal/model"
	"net/http"
)

type IUserRepository interface {
	Create(user *model.User) (uint, error)
	Get(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(email string) (model.User, error)
}

type UserController struct {
	userRepo IUserRepository
}

func NewUserController(userRepo IUserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var userData model.User
	if err := c.BindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := ctrl.userRepo.Create(&userData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"uid": uid})
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	var email struct {
		Email string `json:"email"`
	}
	if err := c.BindJSON(&email); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData, err := ctrl.userRepo.Get(email.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userData})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	var userData model.UserUpdate
	if err := c.BindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Info(userData)

	if err := ctrl.userRepo.Update(&model.User{Email: userData.Email, Password: userData.NewPassword}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userData})
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	var email struct {
		Email string `json:"email"`
	}
	if err := c.BindJSON(&email); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData, err := ctrl.userRepo.Delete(email.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userData})
}
