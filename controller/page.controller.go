package controller

import (
	"strconv"
	"go-decorator/model"
	"go-decorator/schema"
	"github.com/gin-gonic/gin"
)

type PageController struct {
	BaseController
}

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
