package model

import (
	"cts-go/database"
	"cts-go/schema"
	"github.com/jinzhu/gorm"
)

type ProjectModel struct {
	tableName           string
	databaseHandler     *gorm.DB
}

func NewProjectModel() ProjectModel {
	return ProjectModel {
		tableName: "cms_projects",
		databaseHandler: database.GetDatabaseHandleInstance(),
	}
}

func (pm ProjectModel) GetProjectDetail(projectId int) (schema.ProjectSchema, error) {
	var detail schema.ProjectSchema

	if error := pm.databaseHandler.Table(pm.tableName).Where("id = ?", projectId).First(&detail).Error; error != nil {
		return detail, error
	}
	return detail, nil
}

type ProjectListItemSchema struct {
	schema.ProjectSchema
	PageCount      int      `gorm:"column:pageCount; type:int; not null" json:"pageCount"`
}

func (pm ProjectModel) GetProjects(name string, pType, state, pageNo, pageSize int) ([]ProjectListItemSchema, int, error) {
	var total int
	var projects []ProjectListItemSchema

	var whereQueryMap map[string]interface{} = map[string]interface{} {
		"state": state,
	}

	if pType != 0 {
		whereQueryMap["type"] = pType
	}

	error := pm.databaseHandler.Table(pm.tableName).
		Select("cms_projects.*, count(cms_pages.id) as pageCount").
		Joins("left join cms_pages on cms_pages.project_id = cms_projects.id").
		Where(whereQueryMap).Where("cms_projects.name like ?", "%" + name + "%").
		Group("cms_projects.id").Count(&total).Offset(pageNo * pageSize).Limit(pageSize).Find(&projects).Error

	if error != nil {
		return projects, total, error
	}
	return projects, total, nil
}


func (pm ProjectModel) UpdateProject(project schema.ProjectSchema) error {
	var targetProject schema.ProjectSchema

	// 创建
	if project.ID == 0 {
		return pm.databaseHandler.Table(pm.tableName).Create(&project).Error
	}

	// 修改
	targetProject.ID = project.ID
	project.ID = 0

	return pm.databaseHandler.Table(pm.tableName).Model(&targetProject).Updates(project).Error
}

func (pm ProjectModel) RemoveProject(projectId int) error {
	return pm.databaseHandler.Table(pm.tableName).Delete(schema.ProjectSchema{}, "id = ?", projectId).Error
}
