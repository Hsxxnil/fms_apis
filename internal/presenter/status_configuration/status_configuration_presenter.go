package status_configuration

import (
	"fms/internal/interactor/manager/status_configuration"
	"net/http"

	"fms/internal/interactor/pkg/util"

	constant "fms/internal/interactor/constants"

	statusConfigurationModel "fms/internal/interactor/models/status_configurations"
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
	Manager status_configuration.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: status_configuration.Init(db),
	}
}

// Create
// @Summary 新增狀態參數
// @description 新增狀態參數
// @Tags status_configuration
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body status_configurations.Create true "新增狀態參數"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/status-configurations [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &statusConfigurationModel.Create{}
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
// @Summary 取得全部狀態參數
// @description 取得全部狀態參數
// @Tags status_configuration
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=status_configurations.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/status-configurations [get]
// @Router /app/v1.0/status-configurations [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &statusConfigurationModel.Fields{}
	input.FleetID = util.PointerString(ctx.MustGet("fleet_id").(string))
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
// @Summary 取得單一狀態參數
// @description 取得單一狀態參數
// @Tags status_configuration
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param id path string true "狀態參數ID"
// @success 200 object code.SuccessfulMessage{body=status_configurations.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/status-configurations/{id} [get]
// @Router /app/v1.0/status-configurations/{id} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	id := ctx.Param("id")
	input := &statusConfigurationModel.Field{}
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
// @Summary 刪除單一狀態參數
// @description 刪除單一狀態參數
// @Tags status_configuration
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param id path string true "狀態參數ID"
// @param * body status_configurations.Update true "更新狀態參數"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/status-configurations/{id} [delete]
func (c *control) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	input := &statusConfigurationModel.Field{}
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
// @Summary 更新單一狀態參數
// @description 更新單一狀態參數
// @Tags status_configuration
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param id path string true "狀態參數ID"
// @param * body status_configurations.Update true "更新狀態參數"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/status-configurations/{id} [patch]
func (c *control) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	input := &statusConfigurationModel.Update{}
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
