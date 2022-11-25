package controller

import (
	"awesomeProject/model"
	"awesomeProject/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPromotionById(c *gin.Context) {
	id := c.Param("id")
	result := repository.GetPromotionById(id)
	var pm model.Promotion

	if result != pm {
		c.IndentedJSON(http.StatusOK, result)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "promotion not found"})
}
