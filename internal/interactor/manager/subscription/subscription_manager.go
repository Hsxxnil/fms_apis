package subscription

import (
	"encoding/json"
	payModel "fms/internal/interactor/models/payments"
	subscriptionModel "fms/internal/interactor/models/subscriptions"
	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
	"fms/internal/router/middleware/payment"
	"gorm.io/gorm"
)

type Manager interface {
	ActionPay(input *subscriptionModel.ActionPay) (int, any)
	Redirect(input *payModel.PayJson) (int, any)
	Check(input *payModel.PayJson) (int, any)
	Query(input *subscriptionModel.Query) (int, any)
}

type manager struct {
}

func Init(db *gorm.DB) Manager {
	return &manager{}
}

func (m *manager) ActionPay(input *subscriptionModel.ActionPay) (int, any) {
	apiBody := &payModel.PayJson{}
	result, err := payment.ActionPay(input)
	if err != nil {
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(resultByte, &apiBody)
	if err != nil {
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, apiBody)
}

func (m *manager) Redirect(input *payModel.PayJson) (int, any) {
	if input.Status == "SUCCESS" {
		return code.Successful, code.GetCodeMessage(code.Successful, "Pay Successful!")
	} else {
		return code.BadRequest, code.GetCodeMessage(code.BadRequest, "Pay Failed!")
	}
}

func (m *manager) Check(input *payModel.PayJson) (int, any) {
	res, err := payment.GetPayResult(input)
	if err != nil {
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	// todo 交易資料處理
	return code.Successful, code.GetCodeMessage(code.Successful, res)
}

func (m *manager) Query(input *subscriptionModel.Query) (int, any) {
	apiBody, err := payment.Query(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, apiBody)
}

// todo 解析Query回應參數api
