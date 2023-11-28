package repository_users

import (
	"users_manager/internal/domain/entities"
	"users_manager/internal/services/database"
)

func FindByPrimaryKey(key string) entities.User {

	conn := database.Conn
	var user entities.User

	conn.First(&user, key)

	return user
}
