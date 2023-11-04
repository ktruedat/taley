package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/ktruedat/taleBack/pkg/database"
	"net/http"
)

func main() {
	db := database.DBConnect()
	router := gin.Default()
	router.Use(corsMiddleware())

	RegisterUserRoutes(router, db)
	RegisterTaleRoutes(router, db)

	err := router.Run(":1337")
	if err != nil {
		slog.Fatalf("Failed to start server: %v", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,Access-Control-Allow-Origin")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
