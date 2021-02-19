package router

import (
	"fmt"
	"cts-go/config"
	"github.com/gin-gonic/gin"
)

type description struct {
	path        string
	method      string
	handlers    []gin.HandlerFunc
}

var routes []description

func init() {
	routes = append(routes, pageRoutes...)
	routes = append(routes, loginRoutes...)
	routes = append(routes, moduleRoutes...)
	routes = append(routes, projectRoutes...)
	routes = append(routes, userRoutes...)
}

func RegisterRouter(engine *gin.Engine) {
	// 静态目录
	if config.ApplicationConfig["static"] != "" {
		engine.Static("static", config.ApplicationConfig["static"])
	}
	// 模板文件
	if config.ApplicationConfig["templateDir"] != "" {
		engine.LoadHTMLGlob(  fmt.Sprintf("%s/**/*", config.ApplicationConfig["templateDir"]))
	}
	// api 接口
	for _, route := range routes {
		engine.Handle(route.method, route.path, route.handlers...)
	}
}
