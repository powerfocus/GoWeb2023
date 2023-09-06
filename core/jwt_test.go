package core

import (
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

var (
	uuid = UUID(50)
)

func TestCreate(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "admin",
		"isAdmin":  true,
	})
	ss, err := token.SignedString([]byte(uuid))
	if err != nil {
		panic(err)
	}
	t.Logf("%v", ss)
}

func TestCreateFromClaims(t *testing.T) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now()),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(uuid))
	if err != nil {
		panic(err)
	}
	t.Logf("%v", ss)
}

func TestRandString(t *testing.T) {
	str := UUID(50)
	t.Logf("%v", str)
}
