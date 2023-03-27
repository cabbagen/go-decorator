package middleware

import "github.com/gin-gonic/gin"

var globalMiddleware []gin.HandlerFunc = []gin.HandlerFunc {
	HandleCorsMiddleware,
	HandlePanicRecover,
	HandleTokenMiddleware,
}

func RegisterMiddleware(app *gin.Engine) {
	app.Use(globalMiddleware...)
}
