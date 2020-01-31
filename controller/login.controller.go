package controller

import (
	"cts-go/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

type HandleLoginParams struct {
	Username          string    `json:"username"`
	Password          string    `json:"password"`
}
func (lc LoginController) HandleLogin(c *gin.Context) {
	var params HandleLoginParams

	if error := c.BindJSON(&params); error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	userInfo, error := model.NewUserModel().CheckUserInfo(params.Username, params.Password)

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	userInfoBytes, error := json.Marshal(userInfo)

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}
	lc.HandleSuccessResponse(c, string(userInfoBytes))
}
