package password

import "github.com/judwhite/argon2"

func HashPassword(password string) (string, error) {
	str, err := argon2.GenerateFromPassword([]byte(password), argon2.Options{})
	if err != nil {
		return "", err
	}

	return str, nil
}

func CompareHashAndPassword(hash, password string) error {
	if err := argon2.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return err
	}

	return nil
}
