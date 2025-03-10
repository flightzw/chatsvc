package jwt

import (
	"context"
	"strconv"
	"time"

	auth "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func NewClaims() jwtv5.Claims {
	return &jwtv5.RegisteredClaims{}
}

func SignedString(method jwtv5.SigningMethod, id int32, name string, expireIn int32, secret interface{}) (string, error) {
	now := time.Now()
	claims := &jwtv5.RegisteredClaims{
		ExpiresAt: jwtv5.NewNumericDate(now.Add(time.Duration(expireIn) * time.Second)),
		IssuedAt:  jwtv5.NewNumericDate(now),
		ID:        strconv.Itoa(int(id)),
		Subject:   name,
	}
	return jwtv5.NewWithClaims(method, claims).SignedString(secret)
}

func GetUserInfo(ctx context.Context) (id int32, name string) {
	claims, ok := GetRegisteredClaims(ctx)
	if !ok {
		return 0, ""
	}
	tmpId, _ := strconv.Atoi(claims.ID)
	return int32(tmpId), claims.Subject
}

func GetRegisteredClaims(ctx context.Context) (*jwtv5.RegisteredClaims, bool) {
	claims, ok := auth.FromContext(ctx)
	if !ok {
		return nil, ok
	}
	rgtClaims, ok := claims.(*jwtv5.RegisteredClaims)
	if !ok {
		return nil, ok
	}
	return rgtClaims, ok
}
