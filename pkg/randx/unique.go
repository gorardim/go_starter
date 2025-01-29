package randx

import (
	"fmt"
	"hash/fnv"
	"time"
)

const orderIdLen = len("23011223353924499737")

// GenUniqueId 生成订单号
func GenUniqueId() string {
	now := time.Now()
	sn := fmt.Sprintf("%s%03d%04d", now.Format("060102150405"), now.Nanosecond()/1e6, Int(9999))
	// hash str to 0 - 9
	new32 := fnv.New32()
	new32.Write([]byte(sn))
	return fmt.Sprintf("%s%d", sn, new32.Sum32()%10)
}

// VerifyUniqueId 验证订单号
func VerifyUniqueId(tradeNo string) bool {
	if len(tradeNo) != orderIdLen {
		return false
	}
	sn := tradeNo[:orderIdLen-1]
	new32 := fnv.New32()
	new32.Write([]byte(sn))
	return tradeNo[orderIdLen-1:] == fmt.Sprintf("%d", new32.Sum32()%10)
}
