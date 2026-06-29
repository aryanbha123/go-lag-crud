package routes

import (
	"crud-go/internal/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup, db *config.DB) {
	RegisterHealthRoutes(api, db)
	RegisterUserRoutes(api, db)
}
