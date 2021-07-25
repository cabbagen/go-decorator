package controller

import (
	"fmt"
	"time"
	"strconv"
	"go-decorator/model"
	"go-decorator/schema"
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
	IsMark      int          `form:"isMark"`
	PageNo      int          `form:"pageNo"`
	PageSize    int          `form:"pageSize"`
}
func (pc ProjectController) HandleGetProjects(c *gin.Context) {
	var params HandleGetProjectsParams = HandleGetProjectsParams {"", 0, 0,0,0, 10}

	if error := c.BindQuery(&params); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	projects, total, error := model.NewProjectModel().GetProjects(params.Name, params.Type, params.State, params.IsMark, params.PageNo, params.PageSize)

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

func (pc ProjectController) HandleCreateProjectByTemplate(c *gin.Context) {
	templateId, error := strconv.Atoi(c.Query("templateId"))

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	project, error := model.NewTemplateModel().GetProjectByTemplateId(templateId)

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	project.ID = 0
	project.Name = fmt.Sprintf("%s的副本", project.Name)
	project.TemplateId = -1
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	if error := model.NewProjectModel().UpdateProject(project); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	pc.HandleSuccessResponse(c, "ok")
}

