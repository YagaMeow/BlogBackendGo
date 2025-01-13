package utils

import (
	"blog-backend/dao/system/request"
	"blog-backend/global"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.YAGAMI_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseCliams request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.YAGAMI_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.YAGAMI_CONFIG.JWT.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseCliams,
		BufferTime: int64(bf / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"YAGAMI"},
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    global.YAGAMI_CONFIG.JWT.Issuer,
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
