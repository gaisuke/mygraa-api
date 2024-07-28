package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(hashed string, plain string) bool {
	hash, pass := []byte(hashed), []byte(plain)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
