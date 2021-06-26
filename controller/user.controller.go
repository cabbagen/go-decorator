package controller

import (
	"strconv"
	"go-decorator/model"
	"go-decorator/schema"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (uc UserController) HandleGetUserInfo(c *gin.Context) {
	userId, error := strconv.Atoi(c.Param("userId"))

	if error != nil {
		uc.HandleFailResponse(c, error)
		return
	}

	info, error := model.NewUserModel().GetUserInfo(userId)

	if error != nil {
		uc.HandleFailResponse(c, error)
		return
	}
	uc.HandleSuccessResponse(c, info)
}

func (uc UserController) HandleUpdateUserInfo(c *gin.Context) {
	var params schema.UserSchema

	if error := c.BindJSON(&params); error != nil {
		uc.HandleFailResponse(c, error)
		return
	}

	if error := model.NewUserModel().UpdateUserInfo(params); error != nil {
		uc.HandleFailResponse(c, error)
		return
	}
	uc.HandleSuccessResponse(c, "操作成功")
}
