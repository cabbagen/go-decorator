package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlePanicRecover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var panicText = fmt.Sprintf("系统错误 - [url]: %s, [method]: %s, [error]: %s", c.Request.URL.String(), c.Request.Method, err)
			c.JSON(http.StatusOK, gin.H { "status": 500, "msg": panicText, "data": nil})
		}
	}()

	c.Next()
}
