package repository_users

import (
	"users_manager/internal/domain/entities"
	"users_manager/internal/services/database"
)

func Create(user entities.User) entities.User {
	conn := database.Conn
	conn.Create(&user)
	return user
}
