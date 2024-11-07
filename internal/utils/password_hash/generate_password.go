package passwordhash

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(pasword string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)
	return string(password)
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
