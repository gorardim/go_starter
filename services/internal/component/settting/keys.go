package setting

import "app/api/model"

const (
	CustomerServiceUrl = "customer_service_url"  // 客服连接
	UsdtDepositFeeRate = "usdt_deposit_fee_rate" // USDT Deposit fee rate

	BSCWithdrawServiceCharge   = "bsc_withdraw_service_charge"   // BSC提现手续费
	TRC20WithdrawServiceCharge = "trc20_withdraw_service_charge" // TRC20提现手续费

	BalanceWithdrawServiceCharge  = "balance_withdraw_service_charge" // 本金钱包的钱提现时扣除1%手续费
	IncomeWithdrawPointCharge     = "income_withdraw_point_charge"    // 收益钱包提现时扣除5%积分
	UsdtChargeServiceFeeRate      = "usdt_charge_service_fee_rate"    // USDT 充值手续费率
	UsdtWithdrawMinAmount         = "usdt_withdraw_min_amount"        // USDT 提现最小金额
	UsdtNetworkLits               = "usdt_network_list"               // USDT 网络列表
	CrrencyExchangeRate           = "currency_exchange_rate"
	DepositCurrencyExchangeRate   = "deposit_currency_exchange_rate"    // 充值货币汇率
	DepositServiceChargeRate      = "deposit_service_charge_rate"       // 充值手续费率
	DurationAndRateInvestContract = "duration_and_rate_invest_contract" // 期限和利率投资合约
	AboutUsLink                   = "about_us_link"                     // 关于我们
	AboutUsLinkV2                 = "about_us_link_v2"                  // 关于我们
	CustomerServiceLink           = "customer_service_link"             // 客户服务链接
	ScoreTransferServiceFeeRate   = "score_transfer_service_fee_rate"   // 积分转让手续费率
	BalanceTransferServiceFeeRate = "balance_transfer_service_fee_rate" // 余额转让手续费率
	IncomeTransferServiceFeeRate  = "income_transfer_service_fee_rate"  // 收益转让手续费率
	InviteH5Url                   = "invite_h5_url"                     // 邀请H5链接
	H5VideoUrl                    = "h5_video_url"                      // H5视频链接
	H5NewsUrl                     = "h5_news_url"                       // H5新闻链接
	H5PostUrl                     = "h5_post_url"                       // H5帖子链接

	TopDiscoveryDataOnHomePage = "top_discovery_data_home_page_config" // 首页发现顶部数据配置 (json) [{"type":"TYPE","id":1},{"type":"TYPE","id":1}] TYPE: POST, TRAVEL, NEWS, VIDEO, id belong to the type (article, news, travel)
	// H5AboutUrl                    = "h5_about_url"                      // H5关于我们链接 will delete AboutUsLink and use it instead
	// withdraw v2
	BalanceWithdrawV2ServiceCharge = "balance_withdraw_v2_service_charge" // 本金钱包的钱提现时扣除1%手续费
)

type UsdtNetwork struct {
	// Name
	Name string `json:"name"`
	// ChainType
	ChainType string `json:"chain_type"`
	// Icon
	Icon string `json:"icon"`
	// Status
	Status bool `json:"status"`
	// Message
	Message model.LangType `json:"message"`
}

type CurrencyExchangeRateItem struct {
	Lang string `json:"lang"`
	// Name
	Name string `json:"name"`
	// Rate
	Rate float64 `json:"rate"`
	// symbol
	Symbol string `json:"symbol"`
}

type DepositCurrencyExchangeRateItem struct {
	// Name
	Name string `json:"name"`
	// Rate
	Rate float64 `json:"rate"`
	// symbol
	Symbol string `json:"symbol"`
	// bank card
	BankCard string `json:"bank_card"`
	// country name
	CountryName model.LangType `json:"country_name"`
	// country code: USA, CHN, ...
	CountryCode string `json:"country_code"`
	//
	Rules []*model.DisplaySettingItem `json:"rules"`
	//
	RuleDesc string `json:"rule_desc"`
}
