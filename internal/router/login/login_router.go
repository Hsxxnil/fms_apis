package login

import (
	present "fms/internal/presenter/login"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	webV1 := router.Group("fms").Group("web").Group("v1.0")
	{
		webV1.POST("login", control.Login)
		webV1.POST("refresh", control.Refresh)
	}

	appV1 := router.Group("fms").Group("app").Group("v1.0")
	{
		appV1.POST("login", control.Login)
		appV1.POST("refresh", control.Refresh)
	}

	return router
}
