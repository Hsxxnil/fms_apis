package gps

import (
	present "fms/internal/presenter/gps"
	"fms/internal/router/middleware"
	"fms/internal/router/middleware/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	webV1 := router.Group("fms").Group("web").Group("v1.0").Group("gps")
	{
		webV1.GET("list", middleware.Verify(), auth.CheckPermission(), control.WebGetByListNoPagination)
		webV1.POST("list/:licensePlate", middleware.Verify(), auth.CheckPermission(), control.WebGetByLicensePlateList)
	}

	appV1 := router.Group("fms").Group("app").Group("v1.0").Group("gps")
	{
		appV1.GET("list", middleware.Verify(), auth.CheckPermission(), control.AppGetByListNoPagination)
		appV1.POST("list/:licensePlate", middleware.Verify(), auth.CheckPermission(), control.AppGetByLicensePlateList)
	}

	return router
}
