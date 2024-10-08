package gps

import (
	"fms/internal/interactor/manager/gps"
	gpsModel "fms/internal/interactor/models/gps"
	"fms/internal/interactor/pkg/util"
	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	AppGetByListNoPagination(ctx *gin.Context)
	AppGetByLicensePlateList(ctx *gin.Context)
	WebGetByListNoPagination(ctx *gin.Context)
	WebGetByLicensePlateList(ctx *gin.Context)
}

type control struct {
	Manager gps.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: gps.Init(db),
	}
}

// AppGetByListNoPagination
// @Summary 取得全部車輛最新狀態
// @description 取得全部車輛最新狀態
// @Tags gps
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=gps.RealTimeList} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /app/v1.0/gps/list [get]
func (c *control) AppGetByListNoPagination(ctx *gin.Context) {
	input := &gpsModel.Field{}
	input.FleetID = ctx.MustGet("fleet_id").(string)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.AppGetByListNoPagination(input)
	ctx.JSON(httpCode, codeMessage)
}

// AppGetByLicensePlateList
// @Summary 取得單一車輛全部狀態
// @description 取得全部車輛狀態
// @Tags gps
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body gps.Filter false "搜尋"
// @param license-plate path string true "車牌號碼"
// @success 200 object code.SuccessfulMessage{body=gps.RealTimeList} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /app/v1.0/gps/list/{license-plate} [post]
func (c *control) AppGetByLicensePlateList(ctx *gin.Context) {
	licensePlate := ctx.Param("licensePlate")
	input := &gpsModel.Field{}
	input.LicensePlate = util.PointerString(licensePlate)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.AppGetByLicensePlateList(input)
	ctx.JSON(httpCode, codeMessage)
}

// WebGetByListNoPagination
// @Summary 取得全部車輛最新狀態
// @description 取得全部車輛最新狀態
// @Tags gps
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=gps.RealTimeList} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/gps/list [get]
func (c *control) WebGetByListNoPagination(ctx *gin.Context) {
	input := &gpsModel.Field{}
	input.FleetID = ctx.MustGet("fleet_id").(string)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.WebGetByListNoPagination(input)
	ctx.JSON(httpCode, codeMessage)
}

// WebGetByLicensePlateList
// @Summary 取得單一車輛全部狀態
// @description 取得全部車輛狀態
// @Tags gps
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body gps.Filter false "搜尋"
// @param license-plate path string true "車牌號碼"
// @success 200 object code.SuccessfulMessage{body=gps.RealTimeList} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/gps/list/{license-plate} [post]
func (c *control) WebGetByLicensePlateList(ctx *gin.Context) {
	licensePlate := ctx.Param("licensePlate")
	input := &gpsModel.Field{}
	input.LicensePlate = util.PointerString(licensePlate)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.WebGetByLicensePlateList(input)
	ctx.JSON(httpCode, codeMessage)
}
