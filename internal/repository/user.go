package repository

import (
	"context"
	"errors"

	"crud-go/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailTaken   = errors.New("email already in use")
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserRepository struct {
	DB *config.DB
}

func NewUserRepository(db *config.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*User, error) {
	var u User
	err := r.DB.Pool.QueryRow(ctx,
		"SELECT id, name, email FROM users WHERE id = $1", id,
	).Scan(&u.ID, &u.Name, &u.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.DB.Pool.Query(ctx,
		"SELECT id, name, email FROM users ORDER BY name",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *UserRepository) Create(ctx context.Context, name, email, password string) (string, error) {
	var id string
	err := r.DB.Pool.QueryRow(ctx,
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		name, email, password,
	).Scan(&id)
	if isUniqueViolation(err) {
		return "", ErrEmailTaken
	}
	return id, err
}

func (r *UserRepository) Update(ctx context.Context, id, name, email string) error {
	tag, err := r.DB.Pool.Exec(ctx,
		"UPDATE users SET name = $1, email = $2 WHERE id = $3",
		name, email, id,
	)
	if isUniqueViolation(err) {
		return ErrEmailTaken
	}
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	tag, err := r.DB.Pool.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrUserNotFound
	}
	return nil
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
