package main

import (
	"golang-jwt-token/config"
	"golang-jwt-token/controller"
	"golang-jwt-token/helper"
	"golang-jwt-token/model"
	"golang-jwt-token/repository"
	"golang-jwt-token/router"
	"golang-jwt-token/service"
	"log"
	"net/http"
	"github.com/go-playground/validator/v10"
)

func main () {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config file", err)
	}
	//database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	//init repository
	usersRepository := repository.NewUsersRepositoryImpl(db)


	// init service
	authenticationService := service.NewAuthenticationServiceImpl(usersRepository, validate)

	//init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)

	routes := router.NewRouter(authenticationController)
	

	server := &http.Server{
		Addr: ":8888",
		Handler: routes,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)

}