package tokenManager

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenManager interface {
	NewJWT(user_id string) (string, error)
	Parse(access_token string) (*string, error)
	NewRefreshToken() (string, error)
}

type tokenManager struct {
	signin_key string
}

func NewTokenManager(signed_key string) TokenManager {
	return &tokenManager{signin_key: signed_key}
}

// generate new token for 1h
func (t *tokenManager) NewJWT(user_id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		Subject:   user_id,
	})

	return token.SignedString([]byte(t.signin_key))
}

// verify token validity
func (t *tokenManager) Parse(access_token string) (*string, error) {
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.InvalidArgument, "unexpected signing method")
		}
		return []byte(t.signin_key), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("cannot get claims from token")
	}

	id := claims["sub"].(string)

	return &id, nil
}

// refresh token validity
func (t *tokenManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%x", b), nil
}
