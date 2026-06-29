package handlers

import (
	"errors"
	"net/http"

	"crud-go/internal/repository"
	"crud-go/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type updateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Get(c *gin.Context) {
	user, err := h.service.Get(c.Request.Context(), c.Param("id"))
	if errors.Is(err, repository.ErrUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Create(c.Request.Context(), req.Name, req.Email, req.Password)
	if errors.Is(err, repository.ErrEmailTaken) {
		c.JSON(http.StatusConflict, gin.H{"error": "email already in use"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Update(c *gin.Context) {
	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Update(c.Request.Context(), c.Param("id"), req.Name, req.Email)
	if errors.Is(err, repository.ErrUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if errors.Is(err, repository.ErrEmailTaken) {
		c.JSON(http.StatusConflict, gin.H{"error": "email already in use"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	err := h.service.Remove(c.Request.Context(), c.Param("id"))
	if errors.Is(err, repository.ErrUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
