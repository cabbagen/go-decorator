package model

import (
	"cts-go/database"
	"cts-go/schema"
	"github.com/jinzhu/gorm"
)

type PageModel struct {
	tableName           string
	databaseHandler     *gorm.DB
}

func NewPageModel() PageModel {
	return PageModel {
		tableName: "cms_pages",
		databaseHandler: database.GetDatabaseHandleInstance(),
	}
}

func (pm PageModel) GetProjectPages(projectId int) ([]schema.PageSchema, error) {
	var pages []schema.PageSchema

	error := pm.databaseHandler.Table(pm.tableName).Where("project_id = ?", projectId).Find(&pages).Error

	if error != nil {
		return pages, error
	}
	return pages, nil
}

func (pm ProjectModel) UpdateProjectPage(pageInfo schema.PageSchema) error {
	if pageInfo.ID == 0 {
		return pm.databaseHandler.Table(pm.tableName).Create(&pageInfo).Error
	}
	return pm.databaseHandler.Table(pm.tableName).Save(&pageInfo).Error
}

func (pm ProjectModel) RemoveProjectPage(pageId int) error {
	return pm.databaseHandler.Table(pm.tableName).Delete(schema.PageSchema{}, "id = ?", pageId).Error
}
