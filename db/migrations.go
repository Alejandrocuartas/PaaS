package db

import (
	"PaaS/models"
)

func ApplyMigrations() {
	PGDB.LogMode(true)

	// MODELS
	PGDB.AutoMigrate(&models.User{})
	PGDB.AutoMigrate(&models.App{})

	PGDB.Model(&models.App{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
