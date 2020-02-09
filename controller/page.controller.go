package controller

import (
	"cts-go/model"
	"cts-go/schema"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PageController struct {
	BaseController
}

// 获取项目页面
func (pc PageController) HandleGetProjectPages(c *gin.Context) {
	projectId, error := strconv.Atoi(c.Param("projectId"))

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	pages, error := model.NewPageModel().GetProjectPages(projectId)

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	pc.HandleSuccessResponse(c, pages)
}

// 创建 或者 更新 项目页面
func (pc PageController) HandleUpdateProjectPage(c *gin.Context) {
	var params schema.PageSchema

	if error := c.BindJSON(&params); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewPageModel().UpdateProjectPage(params); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	pc.HandleSuccessResponse(c, "操作成功")
}

// 删除项目页面
func (pc PageController) HandleRemoveProjectPage(c *gin.Context) {
	pageId, error := strconv.Atoi(c.Param("pageId"))

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewPageModel().RemoveProjectPage(pageId); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	pc.HandleSuccessResponse(c, "删除成功")
}
