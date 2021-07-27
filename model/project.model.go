package model

import (
	"fmt"
	"go-decorator/schema"
	"go-decorator/database"
	"github.com/jinzhu/gorm"
	"time"
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

func (pm ProjectModel) GetProjects(name string, pType, state, IsMark, pageNo, pageSize int) ([]ProjectListItemSchema, int, error) {
	var total int
	var projects []ProjectListItemSchema
	var whereQueryMap map[string]interface{} = make(map[string]interface{})

	if IsMark != 0 {
		whereQueryMap["is_mark"] = IsMark
	}
	if state != 0 {
		whereQueryMap["state"] = state
	}
	if pType != 0 {
		whereQueryMap["type"] = pType
	}

	error := pm.databaseHandler.Table(pm.tableName).
		Select("cms_projects.*, count(cms_pages.id) as pageCount").
		Joins("left join cms_pages on cms_pages.project_id = cms_projects.id").
		Where("cms_projects.template_id = ?", 0).
		Where(whereQueryMap).
		Where("cms_projects.name like ?", "%" + name + "%").
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

func (pm ProjectModel) CopyProject(id int) (schema.ProjectSchema, error) {
	var pages []schema.PageSchema
	var modules []schema.ModuleSchema

	if error := pm.databaseHandler.Table("cms_pages").Where("project_id = ?", id).Find(&pages).Error; error != nil {
		return schema.ProjectSchema{}, error
	}

	var pageIds []int

	for _, page := range pages {
		pageIds = append(pageIds, page.ID)
	}

	if error := pm.databaseHandler.Table("cms_modules").Where("page_id in (?)", pageIds).Find(&modules).Debug().Error; error != nil {
		return schema.ProjectSchema{}, error
	}

	var moduleInfo map[schema.PageSchema][]schema.ModuleSchema = make(map[schema.PageSchema][]schema.ModuleSchema)

	for _, module := range modules {
		for _, page := range pages {
			if module.PageId == page.ID {
				page.ID = 0
				module.ID = 0
				moduleInfo[page] = append(moduleInfo[page], module)
			}
		}
	}

	var project schema.ProjectSchema

	if error := pm.databaseHandler.Table("cms_projects").First(&project, "id = ?", id).Error; error != nil {
		return schema.ProjectSchema{}, error
	}

	project.ID = 0
	project.TemplateId = 0
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	tx := pm.databaseHandler.Begin()

	if error := tx.Table(pm.tableName).Create(&project).First(&project).Error; error != nil {
		tx.Rollback()
		return schema.ProjectSchema{}, error
	}

	var rawSql string = "insert into cms_modules(page_id, type, sort_no, content, remark) values "

	for page, modules := range moduleInfo {
		page.ProjectId = project.ID
		page.CreatedAt = time.Now()
		page.UpdatedAt = time.Now()

		if error := tx.Table("cms_pages").Create(&page).First(&page).Error; error != nil {
			tx.Rollback()
			return schema.ProjectSchema{}, error
		}

		for _, module := range modules {
			rawSql += fmt.Sprintf("(%d, '%s', %d, '%s', '%s'),", page.ID, module.Type, module.SortNo, module.Content, module.Remark)
		}

		if error := tx.Exec(fmt.Sprintf("%s;", rawSql[0: len(rawSql) - 1])).Debug().Error; error != nil {
			tx.Rollback()
			return schema.ProjectSchema{}, error
		}
	}

	return project, tx.Commit().Error
}
