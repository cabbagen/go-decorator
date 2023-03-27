package model

import (
	"go-decorator/schema"
	"go-decorator/database"
	"github.com/jinzhu/gorm"
)

type PatternModel struct {
	tableName           string
	databaseHandler     *gorm.DB
}

func NewPatternModel() PatternModel {
	return PatternModel {
		tableName: "cms_patterns",
		databaseHandler: database.GetDatabaseHandleInstance(),
	}
}

func (pm PatternModel) GetPatterns() ([]schema.PatternSchema, error) {
	var patterns []schema.PatternSchema

	if error := pm.databaseHandler.Table(pm.tableName).Find(&patterns).Error; error != nil {
		return patterns, error
	}
	return patterns, nil
}
