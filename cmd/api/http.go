package api

import (
	"fmt"
	"jevvonn/bcc-be-freepass-2025/internal/config"
	"jevvonn/bcc-be-freepass-2025/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer() {
	config := config.GetConfig()
	router := gin.Default()

	host := config.GetString("host")
	port := config.GetString("port")

	database.NewDatabase()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(fmt.Sprintf("%s:%s", host, port))
}
