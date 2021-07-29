package controller

import (
	"go-decorator/model"
	"github.com/gin-gonic/gin"
	"go-decorator/schema"
	"strconv"
)

type TemplateController struct {
	BaseController
}

func (tc TemplateController) HandleGetTemplateCategoriesByName(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	templateCategories, error := model.NewTemplateCategoryModel().GetTemplateCategoriesByName(search)

	if error != nil {
		tc.HandleFailResponse(c, error)
		return
	}
	tc.HandleSuccessResponse(c, templateCategories)
}

func (tc TemplateController) HandleSaveTemplate(c *gin.Context) {
	var template schema.TemplateSchema

	if error := c.BindJSON(&template); error != nil {
		tc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewTemplateModel().UpdateTemplate(template); error != nil {
		tc.HandleFailResponse(c, error)
		return
	}

	tc.HandleSuccessResponse(c, "ok");
}

type HandleGetTemplateParams struct {
	Search      string       `form:"search"`
	PageNo      int          `form:"pageNo"`
	PageSize    int          `form:"pageSize"`
}
func (tc TemplateController) HandleGetTemplates(c *gin.Context) {
	params := HandleGetTemplateParams { "", 0, 10 }

	if error := c.BindQuery(&params); error != nil {
		tc.HandleFailResponse(c, error)
		return
	}

	templates, total, error := model.NewTemplateModel().GetTemplates(params.Search, params.PageNo, params.PageSize)

	if error != nil {
		tc.HandleFailResponse(c, error)
		return
	}

	tc.HandleSuccessResponse(c, map[string]interface{}{ "templates": templates, "total": total })
}

func (tc TemplateController) HandleRemoveTemplate(c *gin.Context) {
	templateIdString := c.Param("templateId")

	templateId, error := strconv.Atoi(templateIdString)

	if error != nil {
		tc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewTemplateModel().RemoveTemplate(templateId); error != nil {
		tc.HandleFailResponse(c, error)
		return
	}
	tc.HandleSuccessResponse(c, "ok")
}

func (tc TemplateController) HandleGetTemplateDetail(c *gin.Context)  {
	templateId, error := strconv.Atoi(c.Param("templateId"))

	if error != nil {
		tc.HandleFailResponse(c, error)
		return
	}

	template, error := model.NewTemplateModel().GetTemplateDetail(templateId)

	if error != nil {
		tc.HandleFailResponse(c, error)
		return
	}

	tc.HandleSuccessResponse(c, template)
	return
}
