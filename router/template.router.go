package router

import (
	"github.com/gin-gonic/gin"
	"go-decorator/controller"
)

var templateController controller.TemplateController

var templateRoutes []description = []description{
	description {
		path: "/cms/template/categories/name",
		method: "GET",
		handlers: []gin.HandlerFunc { templateController.HandleGetTemplateCategoriesByName },
	},
	description {
		path: "/cms/template/save",
		method: "POST",
		handlers: []gin.HandlerFunc { templateController.HandleSaveTemplate },
	},
	description {
		path: "/cms/template/search",
		method: "GET",
		handlers: []gin.HandlerFunc { templateController.HandleGetTemplates },
	},
	description {
		path: "/cms/template/remove",
		method: "GET",
		handlers: []gin.HandlerFunc { templateController.HandleRemoveTemplate },
	},
	description {
		path: "/cms/template/:templateId",
		method: "DELETE",
		handlers: []gin.HandlerFunc { templateController.HandleGetTemplateDetail },
	},
}
