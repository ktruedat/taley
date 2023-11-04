package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ktruedat/taleBack/internal/controller"
	"github.com/ktruedat/taleBack/internal/repository"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userCtrl := controller.NewUserController(userRepo)
	route := r.Group("/user")
	route.POST("", userCtrl.CreateUser)
	route.GET("", userCtrl.GetUser)
	route.POST("/update", userCtrl.UpdateUser)
	route.DELETE("", userCtrl.DeleteUser)
}

func RegisterTaleRoutes(r *gin.Engine, db *gorm.DB) {
	taleRepo := repository.NewTaleRepository(db)
	userRepo := repository.NewUserRepository(db)
	taleCtrl := controller.NewTaleController(taleRepo, userRepo)
	originalCtrl := controller.NewOriginalController(db)
	route := r.Group("/tale")
	route.POST("", taleCtrl.CreateTale)
	route.GET("/all", taleCtrl.GetAllTales)
	route.POST("/update", taleCtrl.UpdateTale)
	route.DELETE("", taleCtrl.DeleteTale)

	route.POST("/original", originalCtrl.PostReferences)
	route.GET("/original", originalCtrl.GetReferences)
}
