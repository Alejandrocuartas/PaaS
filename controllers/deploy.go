package controllers

import (
	"PaaS/models"
	"PaaS/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Deploy(c *gin.Context) {

	var err error

	var deployRequest models.Deploy
	if err := c.ShouldBindJSON(&deployRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deployRequest.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.Deploy(deployRequest.AppIdentifier)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "deployment created",
	})
}
