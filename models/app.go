package models

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

const (
	AppStatusActive   = "ACTIVE"
	AppStatusInactive = "INACTIVE"
	AppStatusPending  = "PENDING"
)

type App struct {
	UUID          uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt     sql.NullTime   `gorm:"column:created_at;"`
	UpdatedAt     sql.NullTime   `gorm:"column:updated_at;"`
	DeletedAt     *sql.NullTime  `gorm:"column:deleted_at;" sql:"index"`
	Name          string         `gorm:"column:name;type:varchar(255);"`
	Status        string         `gorm:"column:status;type:varchar(255);"`
	RepositoryUrl string         `gorm:"column:repository_url;type:varchar(255);"`
	DeployUrl     sql.NullString `gorm:"column:deploy_url;type:varchar(255);"`

	//Relations
	UserId uint `gorm:"column:user_id;"`
}

func (u *App) TableName() string {
	return "apps"
}
