package controller

import (
	"fmt"
	"bytes"
	"strconv"
	"archive/zip"
	"cts-go/model"
	"html/template"
	"cts-go/schema"
	"cts-go/resource"
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

type WillDownloadTemplateData struct {
	Data        *[]schema.PageModule
	Title       string
	PageId      int
}
func (pc ProjectController) HandleDownloadProject(c *gin.Context) {
	projectId, error := strconv.Atoi(c.Param("projectId"))

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	modules, error := model.NewModuleModel().GetProjectModules(projectId)

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "源文件.zip"))

	zipFile, templateDataMap := zip.NewWriter(c.Writer), pc.HandlePageModuleAdaptor(modules)

	if error := pc.HandleCreatePageModuleZip(zipFile, templateDataMap); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	if zipErr := zipFile.Close(); zipErr != nil {
		pc.HandleFailResponse(c, error)
	}
}

func (pc ProjectController) HandlePageModuleAdaptor(modules []schema.PageModule) map[int]WillDownloadTemplateData {
	var templateDataMap map[int]WillDownloadTemplateData = make(map[int]WillDownloadTemplateData)
	for _, value := range modules {
		if _, has := templateDataMap[value.PageId]; !has {
			templateDataMap[value.PageId] = WillDownloadTemplateData {
				Data: &[]schema.PageModule { value },
				Title: value.PageName,
				PageId: value.PageId,
			}
			continue
		}
		*templateDataMap[value.PageId].Data = append(*templateDataMap[value.PageId].Data, value)
	}
	return templateDataMap
}

func (pc ProjectController) HandleCreatePageModuleZip(zipFile *zip.Writer, templateDataMap map[int]WillDownloadTemplateData) error {
	resourceTemplate, error := template.New("web_page").Parse(resource.MobileTpl)

	if error != nil {
		return error
	}

	for _, value := range templateDataMap {
		fileBuffer := bytes.NewBuffer([]byte{})

		if error := resourceTemplate.Execute(fileBuffer, value); error != nil {
			return error
		}

		file, fileErr := zipFile.Create(fmt.Sprintf("%s.html", value.Title))

		if fileErr != nil {
			return fileErr
		}

		if _, writeErr := file.Write(fileBuffer.Bytes()); writeErr != nil {
			return writeErr
		}
	}
	return nil
}
