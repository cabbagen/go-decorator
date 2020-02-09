package controller

import (
	"cts-go/model"
	"cts-go/schema"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ModuleController struct {
	BaseController
}

// 获取页面模块
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

// 创建 或者 更新页面模块
func (mc ModuleController) HandleUpdatePageModule(c *gin.Context) {
	var params schema.ModuleSchema

	if error := c.BindJSON(&params); error != nil {
		mc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewModuleModel().UpdatePageModule(params); error != nil {
		mc.HandleFailResponse(c, error)
		return
	}
	mc.HandleSuccessResponse(c, "操作成功")
}

// 删除页面模块
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

// 页面模块排序
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
