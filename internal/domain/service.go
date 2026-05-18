package domain

type UserService interface {
	Create(username, base64EncodedPassword string) (string, error)
	Login(username, base64EncodedPassword string) (string, error)
}
