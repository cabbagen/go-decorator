package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
type BaseController struct {
}

func (bc BaseController) HandleSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H { "status": 200, "msg": nil, "data": data })
}

func (bc BaseController) HandleFailResponse(c *gin.Context, error error) {
	c.JSON(http.StatusOK, gin.H { "status": 500, "msg": error.Error(), "data": nil })
}

func (bc BaseController) HandleFileUpLoad(c *gin.Context) {
	file, error := c.FormFile("file")

	if error != nil {
		bc.HandleFailResponse(c, error)
		return
	}

	dstFilename := fmt.Sprintf("./public/files/%s", file.Filename)

	if error := c.SaveUploadedFile(file, dstFilename); error != nil {
		bc.HandleFailResponse(c, error)
		return
	}
	bc.HandleSuccessResponse(c, fmt.Sprintf("static/files/%s", file.Filename))
}
