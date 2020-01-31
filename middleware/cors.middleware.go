package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var defaultCorsOptions map[string]string = map[string]string {
	"Access-Control-Allow-Origin": "*",
	"Access-Control-Allow-Headers": "content-type, token, app_key",
}

func HandleCorsMiddleware(c *gin.Context) {
	for key, value := range defaultCorsOptions {
		c.Header(key, value)
	}

	if c.Request.Method == "OPTIONS" {
		c.String(http.StatusOK, "true")
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}
