package policy

import (
	present "fms/internal/presenter/policy"
	"fms/internal/router/middleware"
	"fms/internal/router/middleware/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init()
	webV1 := router.Group("fms").Group("web").Group("v1.0").Group("policies")
	{
		webV1.POST("", middleware.Verify(), auth.CheckPermission(), control.Create)
		webV1.GET("", middleware.Verify(), auth.CheckPermission(), control.GetByList)
		webV1.DELETE("", middleware.Verify(), auth.CheckPermission(), control.Delete)
	}

	appV1 := router.Group("fms").Group("app").Group("v1.0").Group("policies")
	{
		appV1.POST("", middleware.Verify(), auth.CheckPermission(), control.Create)
		appV1.GET("", middleware.Verify(), auth.CheckPermission(), control.GetByList)
		appV1.DELETE("", middleware.Verify(), auth.CheckPermission(), control.Delete)
	}

	return router
}
