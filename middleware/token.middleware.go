package middleware

import (
	"strings"
	"net/http"
	"go-decorator/provider"
	"github.com/gin-gonic/gin"
)

var whitelist []string = []string {
	"/cms/login",
	"/cms/captcha",
}

func isInWhitelistUrl(url string) bool {
	for _, value := range whitelist {
		if strings.HasSuffix(url, value) || strings.HasPrefix(url, "/static") {
			return true
		}
	}
	return false
}
func HandleTokenMiddleware(c *gin.Context) {
	if isInWhitelistUrl(c.Request.URL.String()) {
		c.Next()
		return
	}

	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusOK, provider.NewMSCoreResponse(provider.MSCoreResponseTypeMap["FAILED"], nil, "token is not found"))
		return
	}

	if tokenString, error := provider.ParseTokenString(c.GetHeader("Authorization")); error != nil {
		c.AbortWithStatusJSON(http.StatusOK, provider.NewMSCoreResponse(provider.MSCoreResponseTypeMap["FAILED"], nil, "token is error"))
		return
	} else {
		c.Set("parsed-token", tokenString)
	}
	c.Next()
}
