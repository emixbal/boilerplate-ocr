package api

import (
	"boilerplate-ocr/api/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	imageRouter := router.Group("/image")
	{
		routes.ImageRouter(imageRouter)
	}

	router.Run(":4500")
}
