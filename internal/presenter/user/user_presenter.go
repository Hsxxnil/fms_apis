package user

import (
	"net/http"

	"fms/internal/interactor/pkg/util"

	constant "fms/internal/interactor/constants"

	"fms/internal/interactor/manager/user"
	userModel "fms/internal/interactor/models/users"
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
	Manager user.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: user.Init(db),
	}
}

// Create
// @Summary 新增使用者
// @description 新增使用者
// @Tags user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body users.Create true "新增使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/users [post]
// @Router /app/v1.0/users [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &userModel.Create{}
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
// @Summary 取得全部使用者
// @description 取得全部使用者
// @Tags user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=users.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/users [get]
// @Router /app/v1.0/users [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &userModel.Fields{}
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
// @Summary 取得單一使用者
// @description 取得單一使用者
// @Tags user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=users.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/users/current-user [get]
// @Router /app/v1.0/users/current-user [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	input := &userModel.Field{}
	input.ID = ctx.MustGet("user_id").(string)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一使用者
// @description 刪除單一使用者
// @Tags user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param id path string true "使用者ID"
// @param * body users.Update true "更新使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/users/{id} [delete]
// @Router /app/v1.0/users/{id} [delete]
func (c *control) Delete(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	id := ctx.Param("id")
	input := &userModel.Field{}
	input.ID = id
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(trx, input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一使用者
// @description 更新單一使用者
// @Tags user
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param id path string true "使用者ID"
// @param * body users.Update true "更新使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/users/current-user [patch]
// @Router /app/v1.0/users/current-user [patch]
func (c *control) Update(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &userModel.Update{}
	input.ID = ctx.MustGet("user_id").(string)
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(trx, input)
	ctx.JSON(httpCode, codeMessage)
}
