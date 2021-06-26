package controller

import (
	"os"
	"fmt"
	"bytes"
	"strconv"
	"io/ioutil"
	"archive/zip"
	"html/template"
	"go-decorator/model"
	"go-decorator/schema"
	"go-decorator/resource"
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

type TemplateData struct {
	Data        *[]schema.PageModule
	Title       string
	Link        string
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
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%v%s", projectId, ".zip"))

	zipFile, templateDataMap := zip.NewWriter(c.Writer), pc.HandlePageModuleAdaptor(modules)

	if error := pc.HandleCreatePageModuleZip(zipFile, templateDataMap); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	if zipErr := zipFile.Close(); zipErr != nil {
		pc.HandleFailResponse(c, error)
	}
}

func (pc ProjectController) HandlePageModuleAdaptor(modules []schema.PageModule) map[int]TemplateData {
	var templateDataMap map[int]TemplateData = make(map[int]TemplateData)
	for _, value := range modules {
		if _, has := templateDataMap[value.PageId]; !has {
			templateDataMap[value.PageId] = TemplateData {
				Data: &[]schema.PageModule { value },
				Title: value.PageName,
				PageId: value.PageId,
				Link: value.PageLink,
			}
			continue
		}
		*templateDataMap[value.PageId].Data = append(*templateDataMap[value.PageId].Data, value)
	}
	return templateDataMap
}

func (pc ProjectController) HandleCreatePageModuleZip(zipFile *zip.Writer, templateDataMap map[int]TemplateData) error {
	resourceTemplate, error := template.New("web_page").Parse(resource.MobileTpl)

	if error != nil {
		return error
	}

	for _, value := range templateDataMap {
		fileBuffer := bytes.NewBuffer([]byte{})

		if error := resourceTemplate.Execute(fileBuffer, value); error != nil {
			return error
		}

		file, fileErr := zipFile.Create(fmt.Sprintf("%s.html", value.Link))

		if fileErr != nil {
			return fileErr
		}

		if _, writeErr := file.Write(fileBuffer.Bytes()); writeErr != nil {
			return writeErr
		}
	}
	return nil
}

func (pc ProjectController) HandleGeneratePreviewFiles(c *gin.Context) {
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

	if error := os.Mkdir(fmt.Sprintf("./public/preview/%v", projectId), os.ModePerm); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	modulesMap := pc.HandlePageModuleAdaptor(modules)

	resourceTemplate, error := template.New("web_page").Parse(resource.MobileTpl)

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	for _, value := range modulesMap {
		fileBuffer := bytes.NewBuffer([]byte{})

		if error := resourceTemplate.Execute(fileBuffer, value); error != nil {
			pc.HandleFailResponse(c, error)
			return
		}
		if error := ioutil.WriteFile(fmt.Sprintf("./public/preview/%v/%s.html", projectId, value.Link), fileBuffer.Bytes(), os.ModePerm); error != nil {
			pc.HandleFailResponse(c, error)
			return
		}
	}
	pc.HandleSuccessResponse(c, "ok")
}

func (pc ProjectController) HandleRemovePreviewFiles(c *gin.Context) {
	projectId, error := strconv.Atoi(c.Param("projectId"))

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	if error := os.RemoveAll(fmt.Sprintf("./public/preview/%v/", projectId)); error != nil {
		pc.HandleFailResponse(c, error)
		return
	}

	pc.HandleSuccessResponse(c, "ok")
}
