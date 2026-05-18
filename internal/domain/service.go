package domain

type UserService interface {
	Create(username, base64EncodedPassword string) (*UserResponse, *TokenPair, error)
	Login(username, base64EncodedPassword string) (string, error)
}
