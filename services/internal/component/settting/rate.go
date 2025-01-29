package setting

import "app/services/internal/component/lang"

var langRateMap = map[string]string{
	lang.Chinese:   "CNY",
	lang.English:   "USD",
	lang.Vietnam:   "VND",
	lang.Mongolian: "MNT",
	lang.ThaiLand:  "THB",
	lang.India:     "INR",
}

func SelectRate(rateList []*CurrencyExchangeRateItem, lang string) *CurrencyExchangeRateItem {
	for _, v := range rateList {
		if v.Name == langRateMap[lang] {
			return v
		}
	}
	return rateList[0]
}
