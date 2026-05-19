package domain

type UserRequest struct {
	Username              string
	Base64EncodedPassword string
}

type UserResponse struct {
	Id       int64
	Username string
}
