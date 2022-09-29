package service

import "time"

//https://github.com/techschool/pcbook-go/blob/master/service/jwt_manager.go

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}
