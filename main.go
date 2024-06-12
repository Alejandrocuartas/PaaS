package main

import (
	"PaaS/db"
	"PaaS/environment"
	"PaaS/routes"
	"log"
)

func main() {
	environment.InitEnv()

	db.InitializePostgres()

	//db.ApplyMigrations()

	r := routes.SetupRouter()
	r.Run(":8080")
	log.Println("Server is running on port 8080")
}
