package model

import (
	"cts-go/database"
	"cts-go/schema"
	"cts-go/utils"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	tableName           string
	databaseHandler     *gorm.DB
}

func NewUserModel() UserModel {
	return UserModel {
		tableName: "cms_users",
		databaseHandler: database.GetDatabaseHandleInstance(),
	}
}

func (um UserModel) CheckUserInfo(username, password string) (schema.UserSchema, error) {
	var userInfo schema.UserSchema

	if error := um.databaseHandler.Table(um.tableName).Where("username = ? and password = ?", username, utils.Md5(password)).First(&userInfo).Error; error != nil {
		return userInfo, error
	}

	return userInfo, nil
}

