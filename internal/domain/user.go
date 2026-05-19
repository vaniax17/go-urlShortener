package domain

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type User struct {
	Id       int64
	Username string
	Password string // Hashed
}
