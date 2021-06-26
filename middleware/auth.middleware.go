package middleware

import (
	"bytes"
	"encoding/json"
	"go-decorator/schema"
	"github.com/gin-gonic/gin"
)

func HandleAuthMiddleware(c *gin.Context) {
	var userInfo schema.UserSchema

	parsedToken := c.GetHeader("parsed-token")

	if parsedToken == "" {
		c.Next()
		return
	}
	if error := json.Unmarshal(bytes.NewBufferString(parsedToken).Bytes(), &userInfo); error == nil {
		c.Set("userInfo", userInfo)
	}

	c.Next()
}

