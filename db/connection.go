package db

import (
	"PaaS/environment"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	PGDB *gorm.DB
)

func InitializePostgres() {
	var err error
	DbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=require password=%s TimeZone=America/New_York",
		environment.DbHost,
		environment.DbPort,
		environment.DbUser,
		environment.DbName,
		environment.DbPassword)
	PGDB, err = gorm.Open(environment.DbDriver, DbUrl)
	if err != nil {
		fmt.Println("Cannot connect to postgres.go database", err)
	} else {
		fmt.Println("Connected to Postgres!")
	}

	PGDB.LogMode(true)
}
