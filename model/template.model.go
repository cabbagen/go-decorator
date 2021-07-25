package model

import (
	"go-decorator/schema"
	"go-decorator/database"
	"github.com/jinzhu/gorm"
)

type TemplateModel struct {
	tableName           string
	databaseHandler     *gorm.DB
}

func NewTemplateModel() TemplateModel {
	return TemplateModel {
		tableName: "cms_templates",
		databaseHandler: database.GetDatabaseHandleInstance(),
	}
}

func (tm TemplateModel) UpdateTemplate(template schema.TemplateSchema) error {
	var targetTemplate schema.TemplateSchema

	// 修改
	if template.ID > 0 {
		targetTemplate.ID = template.ID
		template.ID = 0

		return tm.databaseHandler.Table(tm.tableName).Model(&targetTemplate).Updates(template).Error
	}

	// 创建
	if error := tm.databaseHandler.Table(tm.tableName).Create(&template).Error; error != nil {
		return error
	}

	targetProject := schema.ProjectSchema {
		TemplateId: template.ID,
	}
	targetProject.ID = template.ProjectId

	return NewProjectModel().UpdateProject(targetProject)
}

func (tm TemplateModel) GetTemplates(search string, pageNo int, pageSize int) ([]schema.TemplateSchema, int, error) {
	var total int
	var templates []schema.TemplateSchema

	if error := tm.databaseHandler.Table(tm.tableName).Where("name LIKE ?", "%" + search + "%").Count(&total).Offset(pageNo * pageSize).Limit(pageSize).Find(&templates).Error; error != nil {
		return templates, total, error
	}
	return templates, total, nil
}

func (tm TemplateModel ) RemoveTemplate(id int) error {
	var template schema.TemplateSchema
	var project schema.ProjectSchema = schema.ProjectSchema { TemplateId: -1 }

	if error := tm.databaseHandler.Table(tm.tableName).Where("id = ?", id).Find(&template).Delete(template).Error; error != nil {
		return error
	}

	project.ID = template.ProjectId

	return NewProjectModel().UpdateProject(project)
}

func (tm TemplateModel) GetTemplateDetail(id int) (schema.TemplateSchema, error) {
	var template schema.TemplateSchema

	if error := tm.databaseHandler.Table(tm.tableName).Find(&template, "id = ?", id).Error; error != nil {
		return template, error
	}
	return template, nil
}

func (tm TemplateModel) GetProjectByTemplateId(id int) (schema.ProjectSchema, error) {
	var project schema.ProjectSchema

	error := tm.databaseHandler.
		Table(tm.tableName).
		Select("cms_projects.*").
		Joins("inner join cms_projects on cms_projects.id = cms_templates.project_id").
		Where("cms_templates.id = ?", id).
		First(&project).
		Error

	if error != nil {
		return project, error
	}
	return project, nil
}
