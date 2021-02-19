package model

import (
	"cts-go/utils"
	"cts-go/schema"
	"cts-go/database"
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

func (um UserModel) GetUserInfo(userId int) (schema.UserSchema, error) {
	var userInfo schema.UserSchema

	if error := um.databaseHandler.Table(um.tableName).Where("id = ?", userId).First(&userInfo).Error; error != nil {
		return userInfo, error
	}
	return userInfo, nil
}

func (um UserModel) UpdateUserInfo(userInfo schema.UserSchema) error {
	var targetProject schema.UserSchema

	// 创建
	if userInfo.ID == 0 {
		return um.databaseHandler.Table(um.tableName).Create(&userInfo).Error
	}

	// 修改
	targetProject.ID = userInfo.ID
	userInfo.ID = 0

	return um.databaseHandler.Table(um.tableName).Model(&targetProject).Updates(userInfo).Error
}
