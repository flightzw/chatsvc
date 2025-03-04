package hash

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGenerateBcryptHashPassword(t *testing.T) {
	type args struct {
		password string
		salt     string
	}
	tests := []args{}
	for i := 0; i < 100; i++ {
		tests = append(tests, args{
			password: GenerateSalt(6 + rand.Intn(10)),
			salt:     GenerateSalt(16),
		})
	}
	for _, tt := range tests {
		hashPassword := GenerateBcryptHashPassword(tt.password, tt.salt)
		fmt.Printf("%-16s %s %s check: %v\n", tt.password, tt.salt, hashPassword, PasswordCheck(hashPassword, tt.password, tt.salt))
	}
}
