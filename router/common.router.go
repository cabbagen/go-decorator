package router

import (
	"cts-go/controller"
	"github.com/gin-gonic/gin"
)

var baseController controller.BaseController

var commonRoutes []description = []description {
	description {
		path: "/handle/upload",
		method: "POST",
		handlers: []gin.HandlerFunc{ baseController.HandleFileUpLoad },
	},
}
