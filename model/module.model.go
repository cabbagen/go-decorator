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

func (mm ModuleModel) UpdateProjectPage(moduleInfo schema.ModuleSchema) error {
	if moduleInfo.ID == 0 {
		return mm.databaseHandler.Table(mm.tableName).Create(&moduleInfo).Error
	}
	return mm.databaseHandler.Table(mm.tableName).Save(&moduleInfo).Error
}

func (mm ModuleModel) RemovePageModule(moduleId int) error {
	return mm.databaseHandler.Table(mm.tableName).Delete(schema.ModuleSchema{}, "id = ?", moduleId).Error
}

func (mm ModuleModel) SortPageModule(pageId int, sortNoInfo []schema.ModuleSort) error {
	var moduleIds []string
	var sql string = fmt.Sprintf("update %s set sort_no = (case id ", mm.tableName)

	for _, value := range sortNoInfo {
		moduleIds = append(moduleIds, string(value.Id))
		sql += fmt.Sprintf("when %d then %d ", value.Id, value.SortNo)
	}

	sql += fmt.Sprintf("end) where id in (%s) and page_id = %d", strings.Join(moduleIds, ","), pageId)

	return mm.databaseHandler.Exec(sql).Error
}

