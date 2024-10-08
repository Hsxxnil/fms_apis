package payments

// PayJson struct is used to action pay for the Payment API.
type PayJson struct {
	// 回傳狀態 回傳SUCCESS:交易付款成功 錯誤代碼:交易付款失敗
	Status string `json:"Status,omitempty"`
	// 商店代號
	MerchantID string `json:"MerchantID,omitempty"`
	// 交易資料 AES加密
	TradeInfo any `json:"TradeInfo,omitempty"`
	// 交易資料 SHA256加密字串
	TradeSha string `json:"TradeSha,omitempty"`
	// 串接程式版本
	Version string `json:"Version,omitempty"`
	// 加密模式 1:加密模式AES/GCM 0或未有此參數:原加密模式AES/CBC/PKCS7Padding
	EncryptType int64 `json:"EncryptType,omitempty"`
}

// QueryJson struct is used to query transaction data for the Payment API.
type QueryJson struct {
	// 商店代號
	MerchantID string `json:"MerchantID,omitempty"`
	// 串接程式版本
	Version string `json:"Version,omitempty"`
	// 回傳格式
	RespondType string `json:"RespondType,omitempty"`
	// 檢查碼
	CheckValue string `json:"CheckValue,omitempty"`
	// 時間戳記(unix)
	TimeStamp int64 `json:"TimeStamp,omitempty"`
	// 商店訂單編號
	MerchantOrderNo string `json:"MerchantOrderNo,omitempty"`
	// 訂單金額 新台幣
	Amt int64 `json:"Amt,omitempty"`
}

// CreditCardJson struct is used to cancel authorization, request payment, refund, cancel request payments, and cancel refunds for the Payment API.
type CreditCardJson struct {
	// 商店代號
	MerchantID string `json:"MerchantID_,omitempty"`
	// 加密資料
	PostData string `json:"PostData_,omitempty"`
}

// UnencryptedData struct represents data before encryption.
type UnencryptedData struct {
	// 商店代號
	MerchantID string `json:"MerchantID,omitempty"`
	// 回傳格式
	RespondType string `json:"RespondType,omitempty"`
	// 時間戳記(unix)
	TimeStamp int64 `json:"TimeStamp,omitempty"`
	// 串接程式版本
	Version string `json:"Version,omitempty"`
	// 語系 en:英文 zh-tw:繁體中文 jp:日文版
	LangType string `json:"LangType,omitempty"`
	// 商店訂單編號 限英、數字、”_ ”格式，長度限制為30字元，同一商店中此編號不可重覆。《取消授權》時與藍新金流交易序號二擇一。
	MerchantOrderNo string `json:"MerchantOrderNo,omitempty"`
	// 訂單金額 新台幣
	Amt int64 `json:"Amt,omitempty"`
	// 商品資訊 限制長度為50字元，編碼為Utf-8，勿使用斷行符號、單引號等特殊符號。
	ItemDesc string `json:"ItemDesc,omitempty"`
	// 交易有效時間 秒數下限為60秒，秒數上限為900秒。 0或未有此參數:不啟用交易限制秒數
	TradeLimit int64 `json:"TradeLimit,omitempty"`
	// 繳費有效期限 僅適用於非即時支付，格式為20140620，最大值為180天。 空值:系統預設7天，自取號時間起算至第7天23:59:59
	ExpireDate string `json:"ExpireDate,omitempty"`
	// 支付完成返回商店網址 只接受80與443 Port
	ReturnURL string `json:"ReturnURL,omitempty"`
	// 支付通知網址 只接受80與443 Port
	NotifyURL string `json:"NotifyURL,omitempty"`
	// 商店取號網址
	CustomerURL string `json:"CustomerURL,omitempty"`
	// 返回商店網址 在藍新支付頁或藍新交易結果頁面上所呈現之返回鈕，返回至指定的頁面。
	ClientBackURL string `json:"ClientBackURL,omitempty"`
	// 付款人電子信箱
	Email string `json:"Email,omitempty"`
	// 藍新金流會員 1:須要登入藍新金流會員 0:不須登入
	LoginType int64 `json:"LoginType,omitempty"`
	// 商店備註 限制長度為300字。
	OrderComment string `json:"OrderComment,omitempty"`
	// 信用卡一次付清
	Credit int64 `json:"CREDIT,omitempty"`
	// GooglePay
	AndroidPay int64 `json:"ANDROIDPAY,omitempty"`
	// SamsungPay
	SamsungPay int64 `json:"SAMSUNGPAY,omitempty"`
	// LINEPay
	LinePay int64 `json:"LINEPAY,omitempty"`
	// LINEPay 產品圖檔連結
	ImageUrl string `json:"ImageUrl,omitempty"`
	// 信用卡分期付款 同時開啟多期別時，將此參數用”,”分隔。 ０或空值:不開啟分期 1:開啟所有分期期別，且不可帶入其他期別參數 3:3期 6:6期 12:12期 18:18期 24:24期 30:30期
	InstFlag string `json:"InstFlag,omitempty"`
	// 信用卡紅利
	CreditRed int64 `json:"CreditRed,omitempty"`
	// 信用卡銀聯卡
	UnionPay int64 `json:"UNIONPAY,omitempty"`
	// 網路ATM 訂單金額超過49,999元或消費者使用手機裝置付款不適用。
	WebATM int64 `json:"WEBATM,omitempty"`
	// ATM轉帳 訂單金額超過49,999元不適用。
	VACC int64 `json:"VACC,omitempty"`
	// 金融機構 若指定1個以上，則用","分隔。 BOT:台灣銀行 HNCB:華南銀行 未帶值:支援所有指定銀行
	BankType string `json:"BankType,omitempty"`
	// 超商代碼繳費 訂單金額小於30元或超過2萬元不適用。
	CVS int64 `json:"CVS,omitempty"`
	// 超商條碼繳費 訂單金額小於20元或超過4萬元不適用。
	BARCODE int64 `json:"BARCODE,omitempty"`
	// 玉山Wallet
	ESUNWallet int64 `json:"ESUNWALLET,omitempty"`
	// 台灣Pay 訂單金額超過49,999元不適用。
	TaiwanPay int64 `json:"TAIWANPAY,omitempty"`
	// 物流 訂單金額大於2萬元不適用，使用前須先登入藍新金流會員專區啟用物流並設定退貨門市與取貨人相關資訊。 1:啟用超商取貨不付款 2: 啟用超商取貨付款 3:啟用超商取貨不付款及超商取貨付款
	CVSCOM int64 `json:"CVSCOM,omitempty"`
	// 簡單付電子錢包
	EzPay int64 `json:"EZPAY,omitempty"`
	// 簡單付微信支付
	EZPWechat int64 `json:"EZPWECHAT,omitempty"`
	// 簡單付支付寶
	EZPAlipay int64 `json:"EZPALIPAY,omitempty"`
	// 物流型態 B2C:大宗寄倉(目前僅支援7-ELEVEN) C2C:店到店(支援7-ELEVEN、全家、萊爾富、OK mart)
	LgsType string `json:"LgsType,omitempty"`
	// 藍新金流交易序號 《取消授權》時與商店訂單編號二擇一。
	TradeNo string `json:"TradeNo,omitempty"`
	// 單號類別 1:使用商店訂單編號 2:使用藍新金流交易單號
	IndexType int64 `json:"IndexType,omitempty"`
	// 請款或退款 1:請款B031/取消請款B033 2:退款B032/取消退款B034
	CloseType int64 `json:"CloseType,omitempty"`
	// 取消請款或退款 1:取消時發動
	Cancel int64 `json:"Cancel,omitempty"`
}

// Response struct represents data after a completed transaction is returned.
type Response struct {
	// 回傳狀態 回傳SUCCESS:交易付款成功 錯誤代碼:交易付款失敗
	Status string `json:"Status,omitempty"`
	// 回傳訊息
	Message string `json:"Message,omitempty"`
	// 回傳參數
	Result *struct {
		// 商店代號
		MerchantID string `json:"MerchantID,omitempty"`
		// 訂單金額 新台幣
		Amt int64 `json:"Amt,omitempty"`
		// 藍新金流交易序號
		TradeNo string `json:"TradeNo,omitempty"`
		// 商店訂單編號
		MerchantOrderNo string `json:"MerchantOrderNo,omitempty"`
		// 支付方式
		PaymentType string `json:"PaymentType,omitempty"`
		// 回傳格式
		RespondType string `json:"RespondType,omitempty"`
		// 支付狀態 0:未付款 1:付款成功 2:付款失敗 3:取消付款 6:退款
		TradeStatus string `json:"TradeStatus,omitempty"`
		// 交易建立時間 回傳格式為：2014-06-25 16:43:49
		CreateTime string `json:"CreateTime,omitempty"`
		// 支付完成時間 回傳格式為：2014-06-25 16:43:49
		PayTime string `json:"PayTime,omitempty"`
		// 檢核碼
		CheckCode string `json:"CheckCode,omitempty"`
		// 預計撥款日 2014-06-25
		FundTime string `json:"FundTime,omitempty"`
		// 實際交易商店代號
		ShopMerchantID string `json:"ShopMerchantID,omitempty"`
		// 交易 IP
		IP string `json:"IP,omitempty"`
		// 款項保管銀行 HNCB:華南銀行 空值:支付寶-玉山銀行、ezPay電子錢包、LINE Pay、使用信用卡支付
		EscrowBank string `json:"EscrowBank,omitempty"`
		// 信用卡支付（一次付清、Google Pay、Samaung Pay、國民旅遊卡、銀聯）
		*Credit
		// WEBATM、ATM繳費
		*ATM
		// 超商代碼繳費
		*CVS
		// 超商條碼繳費
		*Barcode
		// 跨境支付 (包含簡單付電子錢包、簡單付微信支付、簡單付支付寶)
		*CrossBorder
		// 玉山Wallet、台灣Pay
		*Other
		// 超商物流
		*Lgs
	} `json:"Result,omitempty"`
}

// Credit struct is data returned through credit card transactions.
type Credit struct {
	// 收單金融機構 Esun:玉山銀行 Taishin:台新銀行 CTBC:中國信託銀行 NCCC:聯合信用卡中心 CathayBK:國泰世華銀行 Citibank = 花旗銀行 UBOT:聯邦銀行 SKBank:新光銀行 Fubon:富邦銀行 FirstBank:第一銀行
	AuthBank string `json:"AuthBank,omitempty"`
	// 金融機構回應碼
	RespondCode string `json:"RespondCode,omitempty"`
	// 授權碼
	Auth string `json:"Auth,omitempty"`
	// 卡號前六碼
	Card6No string `json:"Card6No,omitempty"`
	// 卡號末四碼
	Card4No string `json:"Card4No,omitempty"`
	// 分期-期別
	Inst int64 `json:"Inst,omitempty"`
	// 分期-首期金額
	InstFirst int64 `json:"InstFirst,omitempty"`
	// 分期-每期金額
	InstEach int64 `json:"InstEach,omitempty"`
	// ECI值 1,2,5,6:3D交易
	ECI string `json:"ECI,omitempty"`
	// 信用卡快速結帳使用狀態 0:該筆交易為非使用信用卡快速結帳功能 1:該筆交易為首次設定信用卡快速結帳功能 2:該筆交易為使用信用卡快速結帳功能 9:該筆交易為取消信用卡快速結帳功能功能
	TokenUseStatus int64 `json:"TokenUseStatus,omitempty"`
	// 紅利折抵後實際金額 僅有使用紅利折抵交易時才會回傳此參數。 0:紅利折抵交易失敗 訂單金額:紅利折抵交易成功
	RedAmt int64 `json:"RedAmt,omitempty"`
	// 交易類別 CREDIT:台灣發卡機構核發之信用卡 FOREIGN:國外發卡機構核發之卡 UNIONPAY:銀聯卡 GOOGLEPAY:GooglePay SAMSUNGPAY:SamsungPay DCC:動態貨幣轉換
	PaymentMethod string `json:"PaymentMethod,omitempty"`
	// 外幣金額
	DCCAmt float64 `json:"DCC_Amt,omitempty"`
	// 匯率
	DCCRate float64 `json:"DCC_Rate,omitempty"`
	// 風險匯率
	DCCMarkup float64 `json:"DCC_Markup,omitempty"`
	// 幣別
	DCCCurrency string `json:"DCC_Currency,omitempty"`
	// 幣別代碼
	DCCCurrencyCode int64 `json:"DCC_Currency_Code,omitempty"`
	// 請款金額
	CloseAmt string `json:"CloseAmt,omitempty"`
	// 請款狀態 0:未請款 1:等待提送請款至收單機構 2:請款處理中 3:請款完成
	CloseStatus string `json:"CloseStatus,omitempty"`
	// 可退款餘額
	BackBalance string `json:"BackBalance,omitempty"`
	// 退款狀態 0:未退款 1:等待提送退款至收單機構 2:退款處理中 3:退款完成
	BackStatus string `json:"BackStatus,omitempty"`
	// 授權結果訊息
	RespondMsg string `json:"RespondMsg,RespondMsg"`
}

// ATM struct is data returned through atm or webAtm transactions.
type ATM struct {
	// 付款人金融機構代碼
	PayBankCode string `json:"PayBankCode,omitempty"`
	// 付款人金融機構帳號末五碼
	PayerAccount5Code string `json:"PayerAccount5Code,omitempty"`
	// 金融機構代碼
	BankCode string `json:"BankCode,omitempty"`
	// 繳費代碼
	CodeNo string `json:"CodeNo,omitempty"`
	// 繳費截止日期 格式為yyyy-mm-dd
	ExpireDate string `json:"ExpireDate,omitempty"`
	// 付款資訊 1.付款方式為超商代碼時，此欄位為超商繳款代碼 2.付款方式為條碼時，此欄位為繳款條碼。此欄位會將三段條碼資訊用逗號”,”組合後回傳 3.付款方式為ATM轉帳時，此欄位為金融機構的轉帳帳號，括號內為金融機構代碼，例：(031)1234567890
	PayInfo string `json:"PayInfo,omitempty"`
	// 交易狀態 0:未付款 1:已付款 2:訂單失敗 3:訂單取消 6:已退款 9:付款中，待銀行確認
	OrderStatus int64 `json:"OrderStatus,omitempty"`
}

// CVS struct is data returned through convenience store transactions.
type CVS struct {
	// 繳費代碼
	CodeNo string `json:"CodeNo,omitempty"`
	// 繳費門市類別 1:7-11 2:全家 3:OK 4:萊爾富
	StoreType int64 `json:"StoreType,omitempty"`
	// 繳費門市代號
	StoreID string `json:"StoreID,omitempty"`
	// 繳費截止日期 格式為yyyy-mm-dd
	ExpireDate string `json:"ExpireDate,omitempty"`
	// 付款資訊 1.付款方式為超商代碼時，此欄位為超商繳款代碼 2.付款方式為條碼時，此欄位為繳款條碼。此欄位會將三段條碼資訊用逗號”,”組合後回傳 3.付款方式為ATM轉帳時，此欄位為金融機構的轉帳帳號，括號內為金融機構代碼，例：(031)1234567890
	PayInfo string `json:"PayInfo,omitempty"`
	// 交易狀態 0:未付款 1:已付款 2:訂單失敗 3:訂單取消 6:已退款 9:付款中，待銀行確認
	OrderStatus int64 `json:"OrderStatus,omitempty"`
}

// Barcode struct is data returned through convenience store barcode transactions.
type Barcode struct {
	// 第一段條碼
	Barcode1 string `json:"Barcode_1,omitempty"`
	// 第二段條碼
	Barcode2 string `json:"Barcode_2,omitempty"`
	// 第三段條碼
	Barcode3 string `json:"Barcode_3,omitempty"`
	// 付款次數
	RepayTimes int64 `json:"RepayTimes,omitempty"`
	// 繳費超商 SEVEN:7-11 FAMILY:全家 OK:OK超商 HILIFE:萊爾富
	PayStore string `json:"PayStore,omitempty"`
	// 繳費截止日期 格式為yyyy-mm-dd
	ExpireDate string `json:"ExpireDate,omitempty"`
	// 付款資訊 1.付款方式為超商代碼時，此欄位為超商繳款代碼 2.付款方式為條碼時，此欄位為繳款條碼。此欄位會將三段條碼資訊用逗號”,”組合後回傳 3.付款方式為ATM轉帳時，此欄位為金融機構的轉帳帳號，括號內為金融機構代碼，例：(031)1234567890
	PayInfo string `json:"PayInfo,omitempty"`
	// 交易狀態 0:未付款 1:已付款 2:訂單失敗 3:訂單取消 6:已退款 9:付款中，待銀行確認
	OrderStatus int64 `json:"OrderStatus,omitempty"`
}

// CrossBorder struct is data returned through cross-border transactions.
type CrossBorder struct {
	// 跨境通路類型 ALIPAY:支付寶 WECHATPAY:微信支付 ACCLINK:約定連結帳戶 CREDIT:信用卡 CVS:超商代碼 P2GEACC:簡單付電子帳戶轉帳 VACC:ATM轉帳 WEBATM:WebATM轉帳
	ChannelID string `json:"ChannelID,omitempty"`
	// 跨境通路交易序號
	ChannelNo string `json:"ChannelNo,omitempty"`
}

// Other struct is data returned through ESUNWallet, TaiwanPay, LinePay transactions.
type Other struct {
	// 實際付款金額
	PayAmt int64 `json:"PayAmt,omitempty"`
	// 紅利折抵金額
	RedDisAmt int64 `json:"RedDisAmt,omitempty"`
	// 繳費截止日期 格式為yyyy-mm-dd
	ExpireDate string `json:"ExpireDate,omitempty"`
	// 付款資訊 1.付款方式為超商代碼時，此欄位為超商繳款代碼 2.付款方式為條碼時，此欄位為繳款條碼。此欄位會將三段條碼資訊用逗號”,”組合後回傳 3.付款方式為ATM轉帳時，此欄位為金融機構的轉帳帳號，括號內為金融機構代碼，例：(031)1234567890
	PayInfo string `json:"PayInfo,omitempty"`
	// 交易狀態 0:未付款 1:已付款 2:訂單失敗 3:訂單取消 6:已退款 9:付款中，待銀行確認
	OrderStatus int64 `json:"OrderStatus,omitempty"`
	// 金融機構回應碼
	RespondCode string `json:"RespondCode,omitempty"`
	// 請款金額
	CloseAmt string `json:"CloseAmt,omitempty"`
	// 請款狀態 0:未請款 1:等待提送請款至收單機構 2:請款處理中 3:請款完成
	CloseStatus string `json:"CloseStatus,omitempty"`
	// 可退款餘額
	BackBalance string `json:"BackBalance,omitempty"`
	// 退款狀態 0:未退款 1:等待提送退款至收單機構 2:退款處理中 3:退款完成
	BackStatus string `json:"BackStatus,omitempty"`
	// 授權結果訊息
	RespondMsg string `json:"RespondMsg,RespondMsg"`
	// 交易類別 CREDIT:台灣發卡機構核發之信用卡 FOREIGN:國外發卡機構核發之卡 UNIONPAY:銀聯卡 GOOGLEPAY:GooglePay SAMSUNGPAY:SamsungPay DCC:動態貨幣轉換
	PaymentMethod string `json:"PaymentMethod,omitempty"`
	// 收單金融機構 Esun:玉山銀行 Taishin:台新銀行 CTBC:中國信託銀行 NCCC:聯合信用卡中心 CathayBK:國泰世華銀行 Citibank = 花旗銀行 UBOT:聯邦銀行 SKBank:新光銀行 Fubon:富邦銀行 FirstBank:第一銀行
	AuthBank string `json:"AuthBank,omitempty"`
}

// Lgs struct is the logistics information returned upon transaction completion.
type Lgs struct {
	// 超商門市編號
	StoreCode string `json:"StoreCode,omitempty"`
	// 超商門市名稱
	StoreName string `json:"StoreName,omitempty"`
	// 超商類別名稱
	StoreType string `json:"StoreType,omitempty"`
	// 超商門市地址
	StoreAddr string `json:"StoreAddr,omitempty"`
	// 取件交易方式 1:取貨付款 3:取貨不付款
	TradeType int64 `json:"TradeType,omitempty"`
	// 取貨人
	CVSCOMName string `json:"CVSCOMName,omitempty"`
	// 取貨人手機號碼
	CVSCOMPhone string `json:"CVSCOMPhone,omitempty"`
	// 物流寄件單號
	LgsNo string `json:"LgsNo,omitempty"`
	// 物流型態
	LgsType string `json:"LgsType,omitempty"`
	// 繳費截止日期 格式為yyyy-mm-dd
	ExpireDate string `json:"ExpireDate,omitempty"`
	// 付款資訊 1.付款方式為超商代碼時，此欄位為超商繳款代碼 2.付款方式為條碼時，此欄位為繳款條碼。此欄位會將三段條碼資訊用逗號”,”組合後回傳 3.付款方式為ATM轉帳時，此欄位為金融機構的轉帳帳號，括號內為金融機構代碼，例：(031)1234567890
	PayInfo string `json:"PayInfo,omitempty"`
	// 交易狀態 0:未付款 1:已付款 2:訂單失敗 3:訂單取消 6:已退款 9:付款中，待銀行確認
	OrderStatus int64 `json:"OrderStatus,omitempty"`
}
