package lang

var Zh = &Lang{
	ErrQueryAppVersion:      "查询版本信息失败",
	ErrQueryHoldAsset:       "查询持有资产失败",
	ErrQueryTotalIncome:     "查询累计收益失败",
	ErrQueryYesterdayIncome: "查询昨日收益失败",
	ErrQuery:                "查询失败",
}

var En = &Lang{
	ErrQueryAppVersion:      "query app version failed",
	ErrQueryHoldAsset:       "query hold asset failed",
	ErrQueryTotalIncome:     "query total income failed",
	ErrQueryYesterdayIncome: "query yesterday income failed",
	ErrQuery:                "query failed",
}

var Mn = &Lang{
	ErrQueryAppVersion:      "асуулга програмын хувилбар амжилтгүй боллоо",
	ErrQueryHoldAsset:       "асуулга барих хөрөнгө амжилтгүй боллоо",
	ErrQueryTotalIncome:     "нийт орлогын асуулга амжилтгүй боллоо",
	ErrQueryYesterdayIncome: "Өчигдрийн орлогын асуулга амжилтгүй боллоо",
	ErrQuery:                "асуулга амжилтгүй болсон",
}

var Vi = &Lang{
	ErrQueryAppVersion:      "Thông tin phiên bản truy vấn không thành công",
	ErrQueryHoldAsset:       "Truy vấn tài sản nắm giữ thất bại",
	ErrQueryTotalIncome:     "Truy vấn thu nhập tích lũy không thành",
	ErrQueryYesterdayIncome: "Kiểm tra thất bại hôm qua",
	ErrQuery:                "Truy vấn không thành công",
}

var I18n = Map{}
