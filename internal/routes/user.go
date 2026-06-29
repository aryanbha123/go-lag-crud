package routes

import (
	"crud-go/internal/config"
	"crud-go/internal/handlers"
	"crud-go/internal/repository"
	"crud-go/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup, db *config.DB) {
	repo := repository.NewUserRepository(db)
	service := services.NewUserService(repo)
	h := handlers.NewUserHandler(service)

	users := api.Group("/users")
	{
		users.GET("", h.List)          // GET    /api/v1/users
		users.GET("/:id", h.Get)       // GET    /api/v1/users/:id
		users.POST("", h.Create)       // POST   /api/v1/users
		users.PUT("/:id", h.Update)    // PUT    /api/v1/users/:id
		users.DELETE("/:id", h.Delete) // DELETE /api/v1/users/:id
	}
}
