package model

import (
	"go-decorator/schema"
	"go-decorator/database"
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

	error := pm.databaseHandler.Table(pm.tableName).Where("project_id = ?", projectId).Order("id").Find(&pages).Error

	if error != nil {
		return pages, error
	}
	return pages, nil
}

func (pm PageModel) UpdateProjectPage(pageInfo schema.PageSchema) error {
	var targetPageInfo schema.PageSchema

	// 创建
	if pageInfo.ID == 0 {
		return pm.databaseHandler.Table(pm.tableName).Create(&pageInfo).Error
	}

	// 修改
	targetPageInfo.ID = pageInfo.ID
	pageInfo.ID = 0

	return pm.databaseHandler.Table(pm.tableName).Model(&targetPageInfo).Updates(pageInfo).Error
}

func (pm PageModel) RemoveProjectPage(pageId int) error {
	return pm.databaseHandler.Table(pm.tableName).Delete(schema.PageSchema{}, "id = ?", pageId).Error
}
