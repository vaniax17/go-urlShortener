package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/vaniax17/go-urlShortener/internal/domain"
)

type userRepo struct {
	*DB
}

func NewUserRepository(db *DB) domain.UserRepository {
	return &userRepo{db}
}

func (r userRepo) Create(ctx context.Context, u *domain.User) (*domain.UserResponse, error) {
	user := User{
		Username: u.Username,
		Password: u.Password,
	}

	if err := r.DB.WithContext(ctx).Create(&user).Error; err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == "23505" {
			return nil, domain.ErrUserAlreadyExists
		}
		return nil, err
	}
	response := domain.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return &response, nil
}

func (r userRepo) Login(ctx context.Context, u *domain.User) (*domain.User, error) {
	user := User{}
	if err := r.DB.WithContext(ctx).First(&user, "username = ?", u.Username).Error; err != nil {
		return nil, err
	}

	u.Password = user.Password

	return u, nil
}
