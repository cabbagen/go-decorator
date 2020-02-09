package model

import (
	"cts-go/database"
	"cts-go/schema"
	"fmt"
	"github.com/jinzhu/gorm"
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

	error := mm.databaseHandler.Table(mm.tableName).Where("page_id = ?", pageId).Find(&modules).Error

	if error != nil {
		return modules, error
	}
	return modules, nil
}

func (mm ModuleModel) UpdatePageModule(moduleInfo schema.ModuleSchema) error {
	var targetModuleInfo schema.ModuleSchema

	if moduleInfo.ID == 0 {
		return mm.databaseHandler.Table(mm.tableName).Create(&moduleInfo).Error
	}

	targetModuleInfo.ID = moduleInfo.ID
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

	return mm.databaseHandler.Exec(sql).Error
}

