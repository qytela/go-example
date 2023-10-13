package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtWrapper struct {
	SecretKey         string
	SecretKeyRefresh  string
	Issuer            string
	ExpirationMinutes int64
	ExpirationHours   int64
}

type JwtClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

func (j *JwtWrapper) GenerateAuthToken(userId uint, secretKey string, expiresAt *jwt.NumericDate) (signedToken string, err error) {
	claims := &JwtClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return
	}

	return
}

func (j *JwtWrapper) GenerateToken(userId uint) (signedToken string, err error) {
	expiresAt := jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)))
	return j.GenerateAuthToken(userId, j.SecretKey, expiresAt)
}

func (j *JwtWrapper) GenerateRefreshToken(userId uint) (signedToken string, err error) {
	expiresAt := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)))
	return j.GenerateAuthToken(userId, j.SecretKeyRefresh, expiresAt)
}

func (j *JwtWrapper) ValidateToken(signedToken string, secretKey string) (claims *JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	return
}
