package service

import (
	"errors"
	"golang-jwt-token/config"
	"golang-jwt-token/data/request"
	"golang-jwt-token/helper"
	"golang-jwt-token/model"
	"golang-jwt-token/repository"
	"golang-jwt-token/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository, 
		Validate: validate,
	}
}

//login auth
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest)(string, error) {
	//find username
	new_user, user_err := a.UsersRepository.FindByUsername(users.Username)
	if user_err != nil {
		return "", errors.New("invalid username or password")
	}

	config, _ := config.LoadConfig(".")
	verify_error := utils.VerifyPassword(new_user.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or password")
	}

	//generate token
	token, err_token := utils.GenerateToken(config.TokenExpireIn, new_user.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil
}


//register auth
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest){
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email: users.Email,
		Password: hashedPassword,
	}

	//save to db
	a.UsersRepository.Save(newUser)

}