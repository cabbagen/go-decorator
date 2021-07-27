package controller

import (
	"strconv"
	"go-decorator/model"
	"go-decorator/schema"
	"github.com/gin-gonic/gin"
	"net/http"
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

	project, error := model.NewTemplateModel().CopyProjectByTemplateId(templateId)

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	pc.HandleSuccessResponse(c, project)
}

func (pc ProjectController) HandleProjectPreview(c *gin.Context) {

	projectId, error := strconv.Atoi(c.Param("projectId"))

	if error != nil {
		c.HTML(http.StatusOK, "error.tmpl", nil)
		return
	}

	project, error := model.NewProjectModel().GetProjectDetail(projectId)

	if error != nil || project.State != 1 {
		c.HTML(http.StatusOK, "error.tmpl", nil)
		return
	}

	pages, error := model.NewPageModel().GetProjectPages(projectId)

	if error != nil {
		c.HTML(http.StatusOK, "error.tmpl", nil)
		return
	}

	var current schema.PageSchema

	for _, page := range pages {
		if c.Param("link") == page.Link[1 : len(page.Link)] {
			current = page
		}
		if c.Param("link") == page.Link[1 : len(page.Link)] + ".html" {
			current = page
		}
	}

	if current.ID == 0 {
		c.HTML(http.StatusOK, "error.tmpl", nil)
		return
	}

	modules, error := model.NewModuleModel().GetPageModules(current.ID)

	if error != nil {
		c.HTML(http.StatusOK, "error.tmpl", nil)
		return
	}

	c.HTML(http.StatusOK, "preview.tmpl", map[string]interface{}{ "Title": current.Name, "Data": modules })
}
