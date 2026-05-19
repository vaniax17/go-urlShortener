package domain

import "context"

type UserService interface {
	Create(ctx context.Context, username, base64EncodedPassword string) (*UserResponse, *TokenPair, error)
	Login(ctx context.Context, username, base64EncodedPassword string) (*UserResponse, *TokenPair, error)
}
