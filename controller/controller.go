package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type BaseController struct {
}

func (bc BaseController) HandleSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H { "status": 200, "msg": nil, "data": data })
}

func (bc BaseController) HandleFailResponse(c *gin.Context, error error) {
	c.JSON(http.StatusOK, gin.H { "status": 500, "msg": error.Error(), "data": nil })
}
