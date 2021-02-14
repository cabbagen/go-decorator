package controller

import (
	"strconv"
	"cts-go/model"
	"cts-go/schema"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	BaseController
}

func (pc ProjectController) HandleGetProjectDetail(c *gin.Context) {
	projectId, error := strconv.Atoi(c.Param("projectId"))

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	info, error := model.NewProjectModel().GetProjectDetail(projectId)

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	pc.HandleSuccessResponse(c, info)
}

type HandleGetProjectsParams struct {
	Name        string       `form:"name"`
	Type        int          `form:"type"`
	State       int          `form:"state"`
	PageNo      int          `form:"pageNo"`
	PageSize    int          `form:"pageSize"`
}
func (pc ProjectController) HandleGetProjects(c *gin.Context) {
	var params HandleGetProjectsParams = HandleGetProjectsParams {"", 0, 1,0, 10}

	if error := c.BindQuery(&params); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	projects, total, error := model.NewProjectModel().GetProjects(params.Name, params.Type, params.State, params.PageNo, params.PageSize)

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	pc.HandleSuccessResponse(c, map[string]interface{} { "projects": projects, "total": total })
}

func (pc ProjectController) HandleUpdateProject(c *gin.Context) {
	var params schema.ProjectSchema

	if error := c.BindJSON(&params); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewProjectModel().UpdateProject(params); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	pc.HandleSuccessResponse(c, "操作成功")
}

func (pc ProjectController) HandleRemoveProject(c *gin.Context) {
	projectId, error := strconv.Atoi(c.Param("projectId"))

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewProjectModel().RemoveProject(projectId); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	pc.HandleSuccessResponse(c, "删除成功")
}
