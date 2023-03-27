package router

import (
	"go-decorator/controller"
	"github.com/gin-gonic/gin"
)

var patternController controller.PatternController

var patternRoutes []description = []description {
	description {
		path: "/cms/patterns",
		method: "GET",
		handlers: []gin.HandlerFunc{ patternController.HandleGetPatterns },
	},
}
