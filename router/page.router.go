package router

import (
	"go-decorator/controller"
	"github.com/gin-gonic/gin"
)

var pageController controller.PageController

var pageRoutes []description = []description {
	description {
		path: "/cms/page/:projectId",
		method: "GET",
		handlers: []gin.HandlerFunc{ pageController.HandleGetProjectPages },
	},
	description {
		path: "/cms/page",
		method: "POST",
		handlers: []gin.HandlerFunc{ pageController.HandleUpdateProjectPage },
	},
	description {
		path: "/cms/page/:pageId",
		method: "DELETE",
		handlers: []gin.HandlerFunc{ pageController.HandleRemoveProjectPage },
	},
}
