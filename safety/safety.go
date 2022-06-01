package safety

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(passwordComHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordComHash), []byte(passwordString))
}
