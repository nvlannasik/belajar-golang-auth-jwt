package service

import "golang-jwt-token/data/request"

type AuthenticationService interface {
	Login(users request.LoginRequest)(string, error)
	Register(users request.CreateUserRequest)
}