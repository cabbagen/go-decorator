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

func (pm ProjectModel) GetProjects(name string, state, pageNo, pageSize int) ([]schema.ProjectSchema, int, error) {
	var total int
	var projects []schema.ProjectSchema

	error := pm.databaseHandler.Table(pm.tableName).Where("state = ? and name like ?", state, "%" + name + "%").Count(&total).Offset(pageNo * pageSize).Limit(pageSize).Find(&projects).Error

	if error != nil {
		return projects, total, error
	}
	return projects, total, nil
}

func (pm ProjectModel) UpdateProject(project schema.ProjectSchema) error {
	if project.ID == 0 {
		return pm.databaseHandler.Table(pm.tableName).Create(&project).Error
	}
	return pm.databaseHandler.Table(pm.tableName).Save(&project).Error
}

func (pm ProjectModel) RemoveProject(projectId int) error {
	return pm.databaseHandler.Table(pm.tableName).Delete(schema.ProjectSchema{}, "id = ?", projectId).Error
}
