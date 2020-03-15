package model

import (
	"cts-go/database"
	"cts-go/schema"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

type ModuleModel struct {
	tableName           string
	databaseHandler     *gorm.DB
}

func NewModuleModel() ModuleModel {
	return ModuleModel {
		tableName: "cms_modules",
		databaseHandler: database.GetDatabaseHandleInstance(),
	}
}

func (mm ModuleModel) GetPageModules(pageId int) ([]schema.ModuleSchema, error) {
	var modules []schema.ModuleSchema

	error := mm.databaseHandler.Table(mm.tableName).Where("page_id = ?", pageId).Order("sort_no").Find(&modules).Error

	if error != nil {
		return modules, error
	}
	return modules, nil
}

func (mm ModuleModel) CreatePageModule(moduleInfo schema.ModuleSchema) error {
	var modules []schema.ModuleSchema
	var sql string = "page_id = ? and sort_no >= ?"

	if error := mm.databaseHandler.Table(mm.tableName).Where(sql, moduleInfo.PageId, moduleInfo.SortNo).Find(&modules).Error; error != nil {
		return error
	}

	if error := mm.databaseHandler.Table(mm.tableName).Create(&moduleInfo).Error; error != nil {
		return error
	}

	if len(modules) == 0 {
		return nil
	}

	var sortInfo []schema.ModuleSort

	for _, module := range modules {
		sortInfo = append(sortInfo, schema.ModuleSort{ Id: module.ID, SortNo: module.SortNo + 1 })
	}

	return mm.SortPageModule(sortInfo)
}

func (mm ModuleModel) UpdatePageModule(moduleInfo schema.ModuleSchema) error {
	var targetModuleInfo schema.ModuleSchema = schema.ModuleSchema {
		BaseSchema: schema.BaseSchema {
			ID: moduleInfo.ID,
		},
	}
	moduleInfo.ID = 0

	return mm.databaseHandler.Table(mm.tableName).Model(&targetModuleInfo).Updates(moduleInfo).Error
}

func (mm ModuleModel) RemovePageModule(moduleId int) error {
	return mm.databaseHandler.Table(mm.tableName).Delete(schema.ModuleSchema{}, "id = ?", moduleId).Error
}

func (mm ModuleModel) SortPageModule(sortNoInfo []schema.ModuleSort) error {
	var moduleIds []string
	var sql string = fmt.Sprintf("update %s set sort_no = (case id ", mm.tableName)

	for _, value := range sortNoInfo {
		moduleIds = append(moduleIds, fmt.Sprintf("%d", value.Id))
		sql += fmt.Sprintf("when %d then %d ", value.Id, value.SortNo)
	}

	sql += fmt.Sprintf("end) where id in (%s)", strings.Join(moduleIds, ","))

	log.Println(sql)

	return mm.databaseHandler.Exec(sql).Error
}
