package util

import (
	"time"

	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})
	return claims.SignedString([]byte(SecretKey))

}

func authSession() func(ctx *fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(SecretKey),
	})
}

func Parsejwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	return claims.Issuer, nil

}
