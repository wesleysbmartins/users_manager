package repository_users

import (
	"users_manager/internal/domain/entities"
	"users_manager/internal/services/database"
)

func FindAll() []entities.User {

	conn := database.Conn
	var users []entities.User
	conn.Find(&users)

	return users
}
