package router

import (
	"golang-jwt-token/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(authenticationController *controller.AuthenticationController) *gin.Engine {
	service := gin.Default()

	service.GET("", func (ctx *gin.Context)  {
		ctx.JSON(200, "Welcome Home")
	})

	router := service.Group("/api/v1")
	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/login", authenticationController.Login)
	authenticationRouter.POST("/register", authenticationController.Register)

	return service



}