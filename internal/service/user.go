package service

import (
	"context"
	"encoding/base64"
	"io"
	"strings"

	"github.com/vaniax17/go-urlShortener/core/jwt"
	"github.com/vaniax17/go-urlShortener/core/password"
	"github.com/vaniax17/go-urlShortener/internal/domain"
)

type userSvc struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userSvc{repo: repo}
}

func (u *userSvc) Create(ctx context.Context, username, base64EncodedPassword string) (*domain.UserResponse, *domain.TokenPair, error) {
	passwordStr, err := decodeBase64(base64EncodedPassword)
	if len(passwordStr) < 9 {
		return nil, nil, domain.ErrPasswordIsTooShort
	}

	if err != nil {
		return nil, nil, err
	}

	hashPassword, err := password.HashPassword(passwordStr)
	if err != nil {
		return nil, nil, err
	}

	user := domain.User{
		Username: username,
		Password: hashPassword,
	}

	createdUser, err := u.repo.Create(ctx, &user)
	if err != nil {
		return nil, nil, err
	}

	tokenPair, err := jwt.GenerateTokenPair(createdUser.Username)
	if err != nil {
		return nil, nil, err
	}

	return createdUser, tokenPair, nil
}

func (u *userSvc) Login(ctx context.Context, username, base64EncodedPassword string) (*domain.UserResponse, *domain.TokenPair, error) {
	passwordStr, err := decodeBase64(base64EncodedPassword)
	if err != nil {
		return nil, nil, err
	}

	user, err := u.repo.Login(ctx, &domain.User{Username: username})
	if err != nil {

	}

	if err := password.CompareHashAndPassword(user.Password, passwordStr); err != nil {
		return nil, nil, err
	}

	tokenPair, err := jwt.GenerateTokenPair(user.Username)
	if err != nil {
		return nil, nil, err
	}

	response := domain.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return &response, tokenPair, nil

}

func decodeBase64(s string) (string, error) {
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(s))
	decoded, err := io.ReadAll(decoder)
	if err != nil {
		return "", nil
	}
	decodedPass := string(decoded)

	return decodedPass, nil
}
