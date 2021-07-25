package model

import (
	"go-decorator/database"
	"github.com/jinzhu/gorm"
	"go-decorator/schema"
)

type TemplateCategoryModel struct {
	tableName           string
	databaseHandler     *gorm.DB
}

func NewTemplateCategoryModel() TemplateCategoryModel {
	return TemplateCategoryModel {
		tableName: "cms_template_category",
		databaseHandler: database.GetDatabaseHandleInstance(),
	}
}

func (tcm TemplateCategoryModel) GetTemplateCategoriesByName(search string) (category []schema.TemplateCategorySchema, err error) {
	var categories []schema.TemplateCategorySchema

	if error := tcm.databaseHandler.Table(tcm.tableName).Where("name like ?","%" + search + "%").Find(&categories).Error; error != nil {
		return categories, error
	}
	return categories, nil
}
