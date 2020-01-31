package router

import (
	"cts-go/controller"
	"github.com/gin-gonic/gin"
)

var loginController controller.LoginController

var loginRoutes []description = []description {
	description {
		path: "/handle/login",
		method: "POST",
		handlers: []gin.HandlerFunc{ loginController.HandleLogin },
	},
}
