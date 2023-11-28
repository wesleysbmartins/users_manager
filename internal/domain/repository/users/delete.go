package repository_users

import (
	"users_manager/internal/domain/entities"
	"users_manager/internal/services/database"
)

func Delete(id string) entities.User {
	conn := database.Conn
	var user entities.User
	conn.Delete(&user, id)
	return user
}
