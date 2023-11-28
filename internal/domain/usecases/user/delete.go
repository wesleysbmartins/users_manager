package usecases

import (
	"users_manager/internal/domain/entities"
	repository_users "users_manager/internal/domain/repository/users"
)

func Delete(id string) entities.User {
	user := repository_users.Delete(id)
	return user
}
