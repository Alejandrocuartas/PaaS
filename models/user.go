package models

import "database/sql"

type User struct {
	ID        uint          `gorm:"primary_key;" sql:"serial; int4"`
	CreatedAt sql.NullTime  `gorm:"column:created_at;"`
	UpdatedAt sql.NullTime  `gorm:"column:updated_at;"`
	DeletedAt *sql.NullTime `gorm:"column:deleted_at;" sql:"index"`
	Name      string        `gorm:"column:name;type:varchar(255);"`
	Email     string        `gorm:"column:email;type:varchar(255);unique_index"`
	Password  string        `gorm:"column:password;type:varchar(255);"`
}

func (u *User) TableName() string {
	return "users"
}
