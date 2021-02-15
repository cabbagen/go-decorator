package router

import (
	"cts-go/controller"
	"github.com/gin-gonic/gin"
)

var projectController controller.ProjectController

var projectRoutes []description = []description {
	description {
		path: "/cms/projects",
		method: "GET",
		handlers: []gin.HandlerFunc{ projectController.HandleGetProjects },
	},
	description {
		path: "/cms/project/:projectId",
		method: "GET",
		handlers: []gin.HandlerFunc{ projectController.HandleGetProjectDetail },
	},
	description {
		path: "/cms/project",
		method: "POST",
		handlers: []gin.HandlerFunc{ projectController.HandleUpdateProject },
	},
	description {
		path: "/cms/project/:projectId",
		method: "DELETE",
		handlers: []gin.HandlerFunc{ projectController.HandleRemoveProject },
	},
	description {
		path: "/cms/projects/download/:projectId",
		method: "GET",
		handlers: []gin.HandlerFunc { projectController.HandleDownloadProject },
	},
}
