package db

import (
	"PaaS/models"
)

func ApplyMigrations() {
	PGDB.LogMode(true)
	PGDB.AutoMigrate(&models.User{})
	PGDB.Model(&models.App{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
