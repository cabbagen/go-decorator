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

	error := pm.databaseHandler.Table(pm.tableName).Where("project_id = ?", projectId).Order("cms_pages.id").Find(&pages).Error

	if error != nil {
		return pages, error
	}
	return pages, nil
}

func (pm PageModel) UpdateProjectPage(pageInfo schema.PageSchema) error {
	var targetPageInfo schema.PageSchema

	// 修改
	if pageInfo.ID > 0 {
		targetPageInfo.ID = pageInfo.ID
		pageInfo.ID = 0

		return pm.databaseHandler.Table(pm.tableName).Model(&targetPageInfo).Updates(pageInfo).Error
	}

	// 新增
	return pm.databaseHandler.Table(pm.tableName).Create(&pageInfo).Find(&pageInfo).Error
}

func (pm PageModel) RemoveProjectPage(pageId int) error {
	return pm.databaseHandler.Table(pm.tableName).Delete(schema.PageSchema{}, "id = ?", pageId).Error
}
