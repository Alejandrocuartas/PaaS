package routes

import (
	"PaaS/controllers"
	"PaaS/environment"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if environment.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.POST("/deploy", controllers.Deploy)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/apps", controllers.CreateApp)
	r.GET("/apps/users/:user_id", controllers.GetApps)

	return r
}
