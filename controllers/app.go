package controllers

import (
	"PaaS/models"
	"PaaS/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateApp(c *gin.Context) {

	var err error

	var data models.CreateApp
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = data.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := services.CreateApp(data)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "app created",
		"data":    r,
	})
}

func GetApps(c *gin.Context) {

	var err error

	//get user_id from params
	userId := c.Param("user_id")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//get apps
	apps, err := services.GetApps(uint(userIdInt))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "apps retrieved",
		"data":    apps,
	})

}
