package domain

type UserRepository interface {
	Create(username, password string) (string, error)
	Login(username, password string) (string, error)
}
