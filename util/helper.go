package util

import (
	//"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/o1egl/paseto"
	//"github.com/fprasty/GoApiWijaya/database"
	//"github.com/fprasty/GoApiWijaya/models"
)

const SecretKey = "udwijaya"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})
	return claims.SignedString([]byte(SecretKey))

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

func GeneratePaseto(issuer string) (string, error) {
	pasetoKey := []byte("UDIWJAYA") // Must be 32 bytes
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		//Audience:   "test",
		Issuer: issuer,
		//Jti:        "123",
		Subject:    "test_subject",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	jsonToken.Set(string(issuer), "")
	footer := "some footer"

	return paseto.NewV2().Encrypt(pasetoKey, jsonToken, footer)
}
