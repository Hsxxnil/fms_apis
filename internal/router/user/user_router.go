package user

import (
	present "fms/internal/presenter/user"
	"fms/internal/router/middleware"
	"fms/internal/router/middleware/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	webV1 := router.Group("fms").Group("web").Group("v1.0").Group("users")
	{
		webV1.POST("", middleware.Verify(), auth.CheckPermission(), middleware.Transaction(db), control.Create)
		webV1.GET("", middleware.Verify(), auth.CheckPermission(), control.GetByList)
		webV1.GET("current-user", middleware.Verify(), auth.CheckPermission(), control.GetBySingle)
		webV1.DELETE(":id", middleware.Verify(), auth.CheckPermission(), middleware.Transaction(db), control.Delete)
		webV1.PATCH("current-user", middleware.Verify(), auth.CheckPermission(), middleware.Transaction(db), control.Update)
	}

	appV1 := router.Group("fms").Group("app").Group("v1.0").Group("users")
	{
		appV1.POST("", middleware.Verify(), auth.CheckPermission(), middleware.Transaction(db), control.Create)
		appV1.GET("", middleware.Verify(), auth.CheckPermission(), control.GetByList)
		appV1.GET("current-user", middleware.Verify(), auth.CheckPermission(), control.GetBySingle)
		appV1.DELETE(":id", middleware.Verify(), auth.CheckPermission(), control.Delete)
		appV1.PATCH("current-user", middleware.Verify(), auth.CheckPermission(), control.Update)
	}

	return router
}
