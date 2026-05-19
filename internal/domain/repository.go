package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, u *User) (*UserResponse, error)
	Login(ctx context.Context, u *User) (*User, error)
}
