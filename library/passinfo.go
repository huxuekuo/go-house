package library

import "github.com/golang-jwt/jwt/v5"

type PassInfo struct {
	UID uint64 `json:"uid"`
	jwt.RegisteredClaims
}
