package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plain string) (hashPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(password, input string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(input))
	if err != nil {
		return err
	}
	return nil
}
