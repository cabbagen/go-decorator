package router

import (
	"cts-go/controller"
	"github.com/gin-gonic/gin"
)

var userController controller.UserController

var userRoutes []description = []description {
	description {
		path: "/cms/user/:userId",
		method: "GET",
		handlers: []gin.HandlerFunc{ userController.HandleGetUserInfo },
	},
	description {
		path: "/cms/user",
		method: "POST",
		handlers: []gin.HandlerFunc{ userController.HandleUpdateUserInfo },
	},
}
