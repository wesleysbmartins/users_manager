package migration

import (
	"users_manager/internal/domain/entities"
	"users_manager/internal/services/database"
)

func Mingrations() {
	conn := database.Conn

	conn.AutoMigrate(&entities.User{})

}
