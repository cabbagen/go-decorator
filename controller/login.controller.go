package controller

import (
	"errors"
	"encoding/json"
	"go-decorator/model"
	"go-decorator/provider"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

// 生成图片验证码
func (lc LoginController) HandleGenerateCaptcha(c *gin.Context) {
	captcha, error := provider.GenerateCaptcha()

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}
	lc.HandleSuccessResponse(c, captcha)
}

// 用户登陆接口
type HandleLoginParams struct {
	Answer            string    `json:"answer"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	CaptchaId         string    `json:"captchaId"`
}
func (lc LoginController) HandleLogin(c *gin.Context) {
	var params HandleLoginParams

	if error := c.BindJSON(&params); error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	// 验证码校验
	if isOk := provider.ValidateCaptcha(params.CaptchaId, params.Answer); !isOk {
		lc.HandleFailResponse(c, errors.New("验证码不正确"))
		return
	}

	// 用户信息校验
	userInfo, error := model.NewUserModel().CheckUserInfo(params.Username, params.Password)

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	// 序列化用户信息
	userInfoBytes, error := json.Marshal(userInfo)

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	// 签发 token
	tokenString, error := provider.SignToken(string(userInfoBytes));

	if error != nil {
		lc.HandleFailResponse(c, error)
		return
	}

	lc.HandleSuccessResponse(c, gin.H{ "token": tokenString, "userInfo": userInfo })
}
