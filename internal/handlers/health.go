package handlers

import (
	"crud-go/internal/config"
	"crud-go/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	DB *config.DB
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	name := c.Query("name")
	response := services.HealthService(name)
	c.JSON(http.StatusOK, response)
}
