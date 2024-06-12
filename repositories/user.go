package repositories

import (
	"PaaS/db"
	"PaaS/models"
)

func CreateUser(user *models.User) error {
	if err := db.PGDB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.PGDB.
		Where("email = ?", email).
		Where("deleted_at IS NULL").
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
