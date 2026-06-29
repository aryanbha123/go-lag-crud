package services

import (
	"context"

	"crud-go/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) List(ctx context.Context) ([]repository.User, error) {
	return s.repo.GetAll(ctx)
}
func (s *UserService) Get(ctx context.Context, id string) (*repository.User, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *UserService) Create(ctx context.Context, name, email, password string) (*repository.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	id, err := s.repo.Create(ctx, name, email, string(hash))
	if err != nil {
		return nil, err
	}
	return &repository.User{ID: id, Name: name, Email: email}, nil
}

func (s *UserService) Update(ctx context.Context, id, name, email string) (*repository.User, error) {
	if err := s.repo.Update(ctx, id, name, email); err != nil {
		return nil, err
	}
	return &repository.User{ID: id, Name: name, Email: email}, nil
}

func (s *UserService) Remove(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
