package router

import (
	"go-decorator/controller"
	"github.com/gin-gonic/gin"
)

var loginController controller.LoginController

var loginRoutes []description = []description {
	description {
		path: "/cms/login",
		method: "POST",
		handlers: []gin.HandlerFunc{ loginController.HandleLogin },
	},
	description {
		path: "/cms/captcha",
		method: "GET",
		handlers: []gin.HandlerFunc { loginController.HandleGenerateCaptcha },
	},
}
