package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ktruedat/taleBack/internal/model"
	"gorm.io/gorm"
	"net/http"
)

type OriginalController struct {
	db *gorm.DB
}

func NewOriginalController(db *gorm.DB) *OriginalController {
	return &OriginalController{db: db}
}

func (ctrl *OriginalController) GetReferences(c *gin.Context) {
	references := &model.References{}
	if err := ctrl.db.Debug().
		Preload("Country").
		Preload("Region").
		Preload("TimeOfTheDay").
		Preload("Weather").
		First(&references).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	//type data struct {
	//	Country      string `json:"country"`
	//	Region       string `json:"region"`
	//	TimeOfTheDay string `json:"timeOfTheDay"`
	//	Weather      string `json:"weather"`
	//}
	//d := data{
	//	Country:      references.Country.Name,
	//	Region:       references.Region.Name,
	//	TimeOfTheDay: references.TimeOfTheDay.Name,
	//	Weather:      references.Weather.Name,
	//}
	c.JSON(http.StatusOK, gin.H{
		"country":      references.Country.Name,
		"region":       references.Region.Name,
		"timeOfTheDay": references.TimeOfTheDay.Name,
		"weather":      references.Weather.Name,
	})
}

func (ctrl *OriginalController) PostReferences(c *gin.Context) {
	var references model.References
	if err := c.BindJSON(&references); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := ctrl.db.Debug().Model(model.References{}).Create(&references).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

//func (ctrl *OriginalController) GetCountry() {
//	//user := &model.User{}
//	//err := repo.dbClient.Debug().Where("email = ?", email).First(&user).Error
//	//if err != nil {
//	//	return nil, err
//	//}
//	//return user, nil
//	country := &model.Country{}
//	if err := ctrl.db.Debug().Where()
//}
//
//func (ctrl *OriginalController) GetRegion() {
//
//}
//
//func (ctrl *OriginalController) GetTimeOfDay() {
//
//}
//
//func (ctrl *OriginalController) GetWeather() {
//
//}
