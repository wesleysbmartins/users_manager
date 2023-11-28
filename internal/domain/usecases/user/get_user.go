package usecases

import (
	"users_manager/internal/domain/entities"
	repository_users "users_manager/internal/domain/repository/users"
)

func GetUser(key string) entities.User {
	user := repository_users.FindByPrimaryKey(key)
	return user
}
