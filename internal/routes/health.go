package routes

import (
	"crud-go/internal/config"
	"crud-go/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(api *gin.RouterGroup, db *config.DB) {
	h := &handlers.HealthHandler{DB: db}
	api.GET("/health", h.HealthCheck)
}
