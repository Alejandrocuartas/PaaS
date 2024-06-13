package repositories

import (
	"PaaS/db"
	"PaaS/models"
)

func CreateApp(app *models.App) error {
	if err := db.PGDB.Create(app).Error; err != nil {
		return err
	}
	return nil
}

func GetApps(userId uint) ([]models.GetAppsResponse, error) {
	var apps []models.GetAppsResponse
	if err := db.PGDB.
		Table("apps").
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}

func UpdateApp(app *models.App) error {
	if err := db.PGDB.Save(app).Error; err != nil {
		return err
	}
	return nil
}

func GetAppByUuid(uuid string) (*models.App, error) {
	var app models.App
	if err := db.PGDB.
		Where("uuid = ?", uuid).
		Where("deleted_at IS NULL").
		First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}
