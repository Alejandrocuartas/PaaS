package routes

import (
	"PaaS/controllers"
	"PaaS/environment"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if environment.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/deploy", controllers.Deploy)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/apps", controllers.CreateApp)
	r.GET("/apps/users/:user_id", controllers.GetApps)

	return r
}
