package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	authToken = "smotriteTokenAhahahahhahahah" // Захардкоденный токен
)

type RequestData struct {
	SomeField string `json:"someField" binding:"required"`
}

// Middleware для проверки авторизации
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != authToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// Обработчик запросов для /endpoint
func endpointHandler(c *gin.Context) {
	var requestData RequestData
	if err := c.ShouldBindWith(&requestData, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "received", "data": requestData})
}

func main() {
	router := gin.Default()

	apiGroup := router.Group("/").Use(AuthMiddleware())
	// POST endpoint
	apiGroup.POST("endpoint", endpointHandler)
	// Запуск
	router.Run(":3000")
}
