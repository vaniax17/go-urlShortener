package domain

type UserRequest struct {
	Username              string
	Base64EncodedPassword string
}

type UserResponse struct {
	Id           uint64
	Username     string
	AccessToken  string
	RefreshToken string
}
