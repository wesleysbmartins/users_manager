package usecases

import (
	"users_manager/internal/domain/entities"
	repository_users "users_manager/internal/domain/repository/users"
)

func GetAllUsers() []entities.User {
	users := repository_users.FindAll()
	return users
}
