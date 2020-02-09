package router

import (
	"cts-go/controller"
	"github.com/gin-gonic/gin"
)

var moduleController controller.ModuleController

var moduleRoutes []description = []description {
	description {
		path: "/cms/module/:pageId",
		method: "GET",
		handlers: []gin.HandlerFunc{ moduleController.HandleGetPageModules },
	},
	description {
		path: "/cms/module",
		method: "POST",
		handlers: []gin.HandlerFunc{ moduleController.HandleUpdatePageModule },
	},
	description {
		path: "/cms/module/sort",
		method: "POST",
		handlers: []gin.HandlerFunc{ moduleController.HandleSortPageModules },
	},
	description {
		path: "/cms/module/:moduleId",
		method: "DELETE",
		handlers: []gin.HandlerFunc{ moduleController.HandleRemovePageModule },
	},
}
