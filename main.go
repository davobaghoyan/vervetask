package main

import (
	"awesomeProject/controller"
	"awesomeProject/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	repository.BulkInsertFromCsv()

	router := gin.Default()
	router.GET("/promotion/:id", controller.GetPromotionById)
	router.Run("localhost:8081")
}
