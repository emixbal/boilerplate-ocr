package controller

import (
	"boilerplate-ocr/api/utils"
	"os"

	"github.com/gin-gonic/gin"
)

/*
 * GET : '/image/extract'
 *
 * @desc Send link image for extract
 *
 * @param  {string} link - Parameters for request
 *
 * @return {object} Request object
 */
func ExtractLink(c *gin.Context) {
	var code int
	var message string
	var result interface{}
	link := c.Request.FormValue("link")

	path, err := utils.SaveImage(link)
	if err != nil {
		code = 500
		message = err.Error()
		result = nil
	}

	extracted := utils.Extract(path)
	data, err := utils.BreakMultiLine(extracted)
	if err != nil {
		code = 500
		message = err.Error()
		result = nil
	}

	err = os.Remove(path)
	if err != nil {
		code = 500
		message = err.Error()
		result = nil
	} else {
		code = 200
		message = "Success extracted file"
		result = data
	}

	response := gin.H{
		"code":    code,
		"message": message,
		"result":  result,
	}

	c.JSON(code, response)
}
