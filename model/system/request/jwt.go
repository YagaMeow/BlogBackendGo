package request

import jwt "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	Id   uint
	Name string
}
