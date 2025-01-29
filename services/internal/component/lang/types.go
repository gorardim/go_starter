package lang

import (
	"context"
	"fmt"

	"app/pkg/ginx"
)

const (
	Chinese   = "zh" // 中文
	English   = "en" // 英文
	Vietnam   = "vi" // 越南
	Mongolian = "mn" // 越南
	ThaiLand  = "th" // 泰国
	India     = "in" // 印度
)

type Map map[string]map[string]string

var GoodsType = map[string]string{"recommended": "recommended", "bestSeller": "bestSeller", "mostRecentUploaded": "mostRecentUploaded"}

func (m Map) Get(local, key string) string {
	if v, ok := m[local]; ok {
		if tran, ok := v[key]; ok {
			return tran
		}
	}

	return key
}

func (m *Map) T(ctx context.Context, key string, args ...interface{}) string {
	value := m.Get(m.local(ctx), key)
	return fmt.Sprintf(value, args...)
}

func (m *Map) local(ctx context.Context) string {
	// 设置了语言
	if trans, ok := TransContext(ctx); ok {
		return trans.Locale()
	}

	// 默认中文
	return Mongolian
}

type Lang struct {
	// 查询版本信息失败
	ErrQueryAppVersion string
	// 查询持有资产失败
	ErrQueryHoldAsset string
	// 查询累计收益失败
	ErrQueryTotalIncome string
	// 查询昨日收益失败
	ErrQueryYesterdayIncome string
	// 查询失败
	ErrQuery string
}

func Get(ctx context.Context) *Lang {
	l := ginx.FromContext(ctx).GetHeader("Accept-Language")
	switch l {
	case English:
		return En
	}
	return Mn
}
