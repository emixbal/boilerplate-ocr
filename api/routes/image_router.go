package routes

import (
	"boilerplate-ocr/api/controller"

	"github.com/gin-gonic/gin"
)

func ImageRouter(router *gin.RouterGroup) {
	router.POST("/extract", controller.ExtractLink)
}
