package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/ktruedat/taleBack/internal/model"
	"net/http"
)

type ITaleRepository interface {
	Create(tale *model.Tale) (uint, error)
	GetAll() ([]model.Tale, error)
	Update(tale *model.Tale) error
	Delete(title string) (model.Tale, error)
}

type IUserGetter interface {
	Get(email string) (*model.User, error)
}

type TaleController struct {
	taleRepo ITaleRepository
	userRepo IUserGetter
}

func NewTaleController(taleRepo ITaleRepository, userRepo IUserGetter) *TaleController {
	return &TaleController{taleRepo: taleRepo, userRepo: userRepo}
}

func (ctrl *TaleController) CreateTale(c *gin.Context) {
	var taleData model.TaleRequest
	if err := c.BindJSON(&taleData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Info(taleData)
	user, err := ctrl.userRepo.Get(taleData.Email)
	if err != nil {
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	tid, err := ctrl.taleRepo.Create(&model.Tale{
		Title:      taleData.Title,
		CategoryID: taleData.Category,
		Contents:   taleData.Contents,
		UserID:     user.ID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tid": tid})
}

func (ctrl *TaleController) GetAllTales(c *gin.Context) {
	tales, err := ctrl.taleRepo.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tales": tales})
}

func (ctrl *TaleController) UpdateTale(c *gin.Context) {
	var taleData model.TaleRequest
	if err := c.BindJSON(&taleData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Info(taleData)
	user, err := ctrl.userRepo.Get(taleData.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.taleRepo.Update(&model.Tale{
		Title:      taleData.Title,
		CategoryID: taleData.Category,
		Contents:   taleData.Contents,
		UserID:     user.ID,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": taleData})
}

func (ctrl *TaleController) DeleteTale(c *gin.Context) {
	var taleTitle struct {
		Title string `json:"title"`
	}
	if err := c.BindJSON(&taleTitle); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taleData, err := ctrl.taleRepo.Delete(taleTitle.Title)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tale deleted": taleData})
}
