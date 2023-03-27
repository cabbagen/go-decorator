package controller

import (
	"go-decorator/model"
	"github.com/gin-gonic/gin"
)

type PatternController struct {
	BaseController
}

func (pc PatternController) HandleGetPatterns(c *gin.Context)  {
	patterns, error := model.NewPatternModel().GetPatterns()

	if error != nil {
		pc.HandleFailResponse(c, error)
		return
	}
	pc.HandleSuccessResponse(c, patterns)
}
