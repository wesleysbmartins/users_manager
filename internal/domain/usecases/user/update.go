package usecases

import (
	"users_manager/internal/domain/entities"
	repository_users "users_manager/internal/domain/repository/users"
)

func Update(user entities.User) entities.User {
	return repository_users.Update(user)
}
