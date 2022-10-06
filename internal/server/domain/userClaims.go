package domain

import "github.com/dgrijalva/jwt-go"

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	jwt.StandardClaims
	Login string `json:"username"`
}
