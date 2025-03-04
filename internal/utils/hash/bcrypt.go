package hash

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateBcryptHashPassword(password, salt string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error generating bcrypt hash: %v", err)
	}
	return base64.StdEncoding.EncodeToString(hashPassword)
}

func PasswordCheck(hashPassword, password, salt string) bool {
	hashPasswordBytes, _ := base64.StdEncoding.DecodeString(hashPassword)
	return bcrypt.CompareHashAndPassword(hashPasswordBytes, []byte(password+salt)) == nil
}
