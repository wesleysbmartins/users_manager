package repository_users

import (
	"users_manager/internal/domain/entities"
	"users_manager/internal/services/database"
)

func Update(user entities.User) entities.User {
	conn := database.Conn

	conn.Save(&user)

	return user
}
