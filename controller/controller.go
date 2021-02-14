package controller

import (
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
