package main

import (
	"jevvonn/bcc-be-freepass-2025/cmd/api"
	"jevvonn/bcc-be-freepass-2025/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.GetConfig()
	gin.SetMode(config.GetString("gin-mode"))

	api.NewHTTPServer()
}
