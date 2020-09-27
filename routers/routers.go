package routers

import (
	"github.com/dacharat/gin-mongodb/apps/welcome"
	"github.com/dacharat/gin-mongodb/config"
	"github.com/gin-gonic/gin"
)

func GenerateRouter() *gin.Engine {
	gin.SetMode(config.GinMode)

	router := gin.New()

	router.GET("/", welcome.WelcomeHandler)

	return router
}
