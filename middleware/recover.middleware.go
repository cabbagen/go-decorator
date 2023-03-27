package middleware

import (
	"fmt"
	"net/http"
	"go-decorator/provider"
	"github.com/gin-gonic/gin"
)

func HandlePanicRecover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var panicText = fmt.Sprintf("系统错误 - [url]: %s, [method]: %s, [error]: %s", c.Request.URL.String(), c.Request.Method, err)
			c.JSON(http.StatusOK, provider.NewMSCoreResponse(provider.MSCoreResponseTypeMap["FAILED"], nil, panicText))
		}
	}()
	c.Next()
}
