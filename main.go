package main

import (
	"cts-go/database"
	"cts-go/middleware"
	"cts-go/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// database
	database.Connect()

	defer database.Destroy()

	// application middleware
	middleware.RegisterMiddleware(engine)

	// application router
	router.RegisterRouter(engine)

	// run
	engine.Run(":7001") // listen and serve on 0.0.0.0:7001 (for windows "localhost:7001")
}
