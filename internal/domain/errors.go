package domain

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrDecodingBase64     = errors.New("can't decode base64 (EOF)")
	ErrPasswordIsTooShort = errors.New("password is too short (min 9)")
)
