package subscriptions

// ActionPay struct is used to encryption subscription data.
type ActionPay struct {
	// 訂單金額 新台幣
	Amount int64 `json:"amount,omitempty"`
	// 商品資訊 限制長度為50字元，編碼為Utf-8，勿使用斷行符號、單引號等特殊符號。
	Description string `json:"description,omitempty"`
	// 付款人電子信箱
	Email string `json:"email,omitempty"`
}

// Query struct is used to query subscription data.
type Query struct {
	// 訂單金額 新台幣
	Amount int64 `json:"amount"`
	// 訂單編號
	MerchantOrderNo string `json:"merchant_order_no"`
}
