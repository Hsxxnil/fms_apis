package subscription

import (
	"fms/internal/interactor/manager/subscription"
	payModel "fms/internal/interactor/models/payments"
	subscriptionModel "fms/internal/interactor/models/subscriptions"
	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
)

type Control interface {
	ActionPay(ctx *gin.Context)
	Redirect(ctx *gin.Context)
	Check(ctx *gin.Context)
	Query(ctx *gin.Context)
}

type control struct {
	Manager subscription.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: subscription.Init(db),
	}
}

// ActionPay
// @Summary 付款
// @description 付款
// @Tags subscription
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body subscriptions.ActionPay true "付款"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/subscriptions/action-pay [post]
func (c *control) ActionPay(ctx *gin.Context) {
	input := &subscriptionModel.ActionPay{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	httpCode, codeMessage := c.Manager.ActionPay(input)
	ctx.JSON(httpCode, codeMessage)
}

// Redirect
// @Summary 供藍新呼叫提供跳轉介面
// @description 供藍新呼叫提供跳轉介面
// @Tags subscription
// @version 1.0
// @Accept json
// @produce json
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/subscriptions/redirect [post]
func (c *control) Redirect(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	// 解析request body
	queryValues, err := url.ParseQuery(string(body))
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	payData := &payModel.PayJson{
		Status:     queryValues.Get("Status"),
		MerchantID: queryValues.Get("MerchantID"),
		TradeInfo:  queryValues.Get("TradeInfo"),
		TradeSha:   queryValues.Get("TradeSha"),
		Version:    queryValues.Get("Version"),
	}

	httpCode, _ := c.Manager.Redirect(payData)
	if httpCode == 200 {
		ctx.Redirect(http.StatusMovedPermanently, "https://d.fms.jinher-net.com/pay_success")
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "https://d.fms.jinher-net.com/pay_failed")
	}
}

// Check
// @Summary 供藍新呼叫確認交易狀態
// @description 供藍新呼叫確認交易狀態
// @Tags subscription
// @version 1.0
// @Accept json
// @produce json
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/subscriptions/check [post]
func (c *control) Check(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	// 解析request body
	queryValues, err := url.ParseQuery(string(body))
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	payData := &payModel.PayJson{
		Status:     queryValues.Get("Status"),
		MerchantID: queryValues.Get("MerchantID"),
		TradeInfo:  queryValues.Get("TradeInfo"),
		TradeSha:   queryValues.Get("TradeSha"),
		Version:    queryValues.Get("Version"),
	}

	httpCode, codeMessage := c.Manager.Check(payData)
	ctx.JSON(httpCode, codeMessage)
}

// Query
// @Summary 單筆交易查詢
// @description 單筆交易查詢
// @Tags subscription
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body subscriptions.Query true "單筆交易查詢"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /web/v1.0/subscriptions/query [post]
func (c *control) Query(ctx *gin.Context) {
	input := &subscriptionModel.Query{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	httpCode, codeMessage := c.Manager.Query(input)
	ctx.JSON(httpCode, codeMessage)
}
