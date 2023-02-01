package repository

import "golang-jwt-token/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(users int)
	FindById(usersId int)(model.Users, error)
	FindAll() []model.Users
	FindByUsername(username string)(model.Users, error)
}