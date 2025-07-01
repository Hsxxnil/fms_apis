package payment

import (
	"encoding/hex"
	"encoding/json"
	"fms/config"
	payModel "fms/internal/interactor/models/payments"
	subscriptionModel "fms/internal/interactor/models/subscriptions"
	"fms/internal/interactor/pkg/util/encryption"
	"fms/internal/interactor/pkg/util/hash"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	returnUrl  = fmt.Sprintf("%s/fms/web/v1.0/subscriptions/redirect", config.Domain)
	notifyUrl  = fmt.Sprintf("%s/fms/web/v1.0/subscriptions/check", config.Domain)
)

// ActionPay is used to newebpay api NPA-F01.
func ActionPay(input *subscriptionModel.ActionPay) (*payModel.PayJson, error) {
	taiwanTimeZone, _ := time.LoadLocation("Asia/Taipei")
	tradeInfo := payModel.UnencryptedData{
		MerchantID:      config.PaymentMid,
		RespondType:     "JSON",
		TimeStamp:       time.Now().In(taiwanTimeZone).Unix(),
		Version:         "2.0",
		ReturnURL:       returnUrl,
		NotifyURL:       notifyUrl,
		Credit:          1,
		MerchantOrderNo: strconv.FormatInt(time.Now().In(taiwanTimeZone).Unix(), 10),
		Amt:             input.Amount,
		ItemDesc:        input.Description,
		Email:           input.Email,
	}

	tradeInfoByte, err := json.Marshal(tradeInfo)
	if err != nil {
		return nil, err
	}

	var tradeInfoMap map[string]any
	err = json.Unmarshal(tradeInfoByte, &tradeInfoMap)
	if err != nil {
		return nil, err
	}

	// 排序
	sortedStr := SortMapToString(tradeInfoMap)

	// 將請求字串使用AES-256-CBC(PKCS7填充)加密
	encrypted, err := encryption.AesEncryptCBC([]byte(sortedStr), []byte(config.PaymentKey), []byte(config.PaymentIv))
	if err != nil {
		return nil, err
	}
	encryptedStr := hex.EncodeToString(encrypted)

	// 將AES加密字串產生檢查碼
	hashes := fmt.Sprintf("HashKey=%s&%s&HashIV=%s", config.PaymentKey, encryptedStr, config.PaymentIv)
	hashStr := strings.ToUpper(hash.Sha256(hashes))

	// 回傳加密資訊
	apiBody := &payModel.PayJson{
		MerchantID: config.PaymentMid,
		Version:    "2.0",
		TradeInfo:  encryptedStr,
		TradeSha:   hashStr,
	}

	return apiBody, nil
}

// Query is used to newebpay api NPA-B02.
func Query(input *subscriptionModel.Query) (*payModel.QueryJson, error) {
	query := payModel.UnencryptedData{
		MerchantID:      config.PaymentMid,
		MerchantOrderNo: input.MerchantOrderNo,
		Amt:             input.Amount,
	}

	queryByte, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	var queryMap map[string]any
	err = json.Unmarshal(queryByte, &queryMap)
	if err != nil {
		return nil, err
	}

	// 排序
	sortedStr := SortMapToString(queryMap)

	// 將AES加密字串產生檢查碼
	hashes := fmt.Sprintf("IV=%s&%s&Key=%s", config.PaymentIv, sortedStr, config.PaymentKey)
	hashStr := strings.ToUpper(hash.Sha256(hashes))

	// 回傳加密資訊
	taiwanTimeZone, _ := time.LoadLocation("Asia/Taipei")
	apiBody := &payModel.QueryJson{
		MerchantID:      config.PaymentMid,
		Version:         "1.3",
		RespondType:     "JSON",
		CheckValue:      hashStr,
		TimeStamp:       time.Now().In(taiwanTimeZone).Unix(),
		MerchantOrderNo: input.MerchantOrderNo,
		Amt:             input.Amount,
	}

	return apiBody, nil
}

// GetPayResult is used to get pay result.
func GetPayResult(input *payModel.PayJson) (*payModel.Response, error) {
	encryptedBytes, err := hex.DecodeString(input.TradeInfo.(string))
	if err != nil {
		return nil, err
	}

	// 使用AES-256-CBC解密
	decryptedBytes, err := encryption.AesDecryptCBC(encryptedBytes, []byte(config.PaymentKey), []byte(config.PaymentIv))
	if err != nil {
		return nil, err
	}

	res := &payModel.Response{}
	err = json.Unmarshal(decryptedBytes, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// SortMapToString is a generic function to sort map to string.
func SortMapToString(input map[string]any) string {
	var str string
	var keys []string
	for mapKey := range input {
		keys = append(keys, mapKey)
	}
	sort.Strings(keys)
	for i, mapKey := range keys {
		if mapKey == "TimeStamp" {
			str += fmt.Sprintf("%s=%d", mapKey, int64(input[mapKey].(float64)))
		} else {
			str += fmt.Sprintf("%s=%v", mapKey, input[mapKey])
		}
		if i < len(keys)-1 {
			str += "&"
		}
	}
	return str
}
