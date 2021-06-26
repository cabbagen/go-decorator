package controller

import (
	"strconv"
	"go-decorator/model"
	"go-decorator/schema"
	"github.com/gin-gonic/gin"
)

type ModuleController struct {
	BaseController
}

func (mc ModuleController) HandleGetPageModules(c *gin.Context) {
	pageId, error := strconv.Atoi(c.Param("pageId"))

	if error != nil {
		mc.HandleFailResponse(c, error)
		return
	}

	modules, error := model.NewModuleModel().GetPageModules(pageId)

	if error != nil {
		mc.HandleFailResponse(c, error)
		return
	}
	mc.HandleSuccessResponse(c, modules)
}

func (mc ModuleController) HandleUpdatePageModule(c *gin.Context) {
	var params schema.ModuleSchema

	if error := c.BindJSON(&params); error != nil {
		mc.HandleFailResponse(c, error)
		return
	}
	if params.ID == 0 {
		mc.handleInnerCreatePageModule(c, params)
		return
	}
	mc.handleInnerUpdatePageModule(c, params)
}

func (mc ModuleController) handleInnerCreatePageModule(c *gin.Context, params schema.ModuleSchema) {
	if error := model.NewModuleModel().CreatePageModule(params); error != nil {
		mc.HandleFailResponse(c, error)
		return
	}
	mc.HandleSuccessResponse(c, "操作成功")
}

func (mc ModuleController) handleInnerUpdatePageModule(c *gin.Context, params schema.ModuleSchema) {
	if error := model.NewModuleModel().UpdatePageModule(params); error != nil {
		mc.HandleFailResponse(c, error)
		return
	}
	mc.HandleSuccessResponse(c, "操作成功")
}

func (mc ModuleController) HandleRemovePageModule(c *gin.Context) {
	moduleId, error := strconv.Atoi(c.Param("moduleId"))

	if error != nil {
		mc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewModuleModel().RemovePageModule(moduleId); error != nil {
		mc.HandleFailResponse(c, error)
		return
	}
	mc.HandleSuccessResponse(c, "删除成功")
}

type HandleSortPageModulesParams struct {
	SortInfo          []schema.ModuleSort        `form:"sortInfo"`
}
func (mc ModuleController) HandleSortPageModules(c *gin.Context) {
	var params HandleSortPageModulesParams

	if error := c.BindJSON(&params); error != nil {
		mc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewModuleModel().SortPageModule(params.SortInfo); error != nil {
		mc.HandleFailResponse(c, error)
		return
	}
	mc.HandleSuccessResponse(c, "操作成功")
}
