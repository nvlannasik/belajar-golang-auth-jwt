package controller

import (
	"golang-jwt-token/data/request"
	"golang-jwt-token/data/response"
	"golang-jwt-token/helper"
	"golang-jwt-token/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		authenticationService: service,
	}
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := controller.authenticationService.Login(loginRequest)

	if err_token != nil {
		webResponse := response.Response{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Message: "Invalid Username or Password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token: token,
	}

	webResponse := response.Response{
		Code: 200,
		Status: "OK",
		Message: "Login Success",
		Data: resp,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.authenticationService.Register(createUserRequest)

	webResponse := response.Response{
		Code: 200,
		Status: "OK",
		Message: "Register Success",
		Data: nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
