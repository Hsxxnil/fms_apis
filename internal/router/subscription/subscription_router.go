package subscription

import (
	present "fms/internal/presenter/subscription"
	"fms/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	control := present.Init(db)
	webV1 := router.Group("fms").Group("web").Group("v1.0").Group("subscriptions")
	{
		webV1.POST("action-pay", middleware.Verify(), control.ActionPay)
		webV1.POST("redirect", control.Redirect)
		webV1.POST("check", control.Check)
		webV1.POST("query", middleware.Verify(), control.Query)
	}

	return router
}
