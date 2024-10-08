package gps_device

import (
	"fms/internal/interactor/manager/gps_device"
	"net/http"

	"fms/internal/interactor/pkg/util"

	constant "fms/internal/interactor/constants"

	gpsDeviceModel "fms/internal/interactor/models/gps_devices"
	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	Create(ctx *gin.Context)
	GetByList(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager gps_device.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: gps_device.Init(db),
	}
}

// Create
// @Summary 新增Gps設備
// @description 新增Gps設備
// @Tags gps_device
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body gps_devices.Create true "新增Gps設備"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/gps-devices [post]
// @Router /app/v1.0/gps-devices [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &gpsDeviceModel.Create{}
	input.CreatedBy = ctx.MustGet("user_id").(string)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Create(trx, input)
	ctx.JSON(httpCode, codeMessage)
}

// GetByList
// @Summary 取得全部Gps設備
// @description 取得全部Gps設備
// @Tags gps_device
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=gps_devices.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/gps-devices [get]
// @Router /app/v1.0/gps-devices [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &gpsDeviceModel.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= constant.DefaultLimit {
		input.Limit = constant.DefaultLimit
	}

	httpCode, codeMessage := c.Manager.GetByList(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingle
// @Summary 取得單一Gps設備
// @description 取得單一Gps設備
// @Tags gps_device
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param id path string true "Gps設備ID"
// @success 200 object code.SuccessfulMessage{body=gps_devices.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/gps-devices/{id} [get]
// @Router /app/v1.0/gps-devices/{id} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	id := ctx.Param("id")
	input := &gpsDeviceModel.Field{}
	input.ID = id
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一Gps設備
// @description 刪除單一Gps設備
// @Tags gps_device
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param id path string true "Gps設備ID"
// @param * body gps_devices.Update true "更新Gps設備"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/gps-devices/{id} [delete]
// @Router /app/v1.0/gps-devices/{id} [delete]
func (c *control) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	input := &gpsDeviceModel.Field{}
	input.ID = id
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一Gps設備
// @description 更新單一Gps設備
// @Tags gps_device
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param id path string true "Gps設備ID"
// @param * body gps_devices.Update true "更新Gps設備"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/gps-devices/{id} [patch]
// @Router /app/v1.0/gps-devices/{id} [patch]
func (c *control) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	input := &gpsDeviceModel.Update{}
	input.ID = id
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
