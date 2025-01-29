package utils

import (
	"fmt"
	"math"
	"strconv"
)

func Fen2Yuan(fee int) string {
	return fmt.Sprintf("%.2f", float64(fee)/100)
}

func Yuan2Fen(yuan string) (int, error) {
	float, err := strconv.ParseFloat(yuan, 64)
	if err != nil {
		return 0, err
	}
	return int(float * 100), nil
}

// FormatAmount 小数点后第三位四舍五入（1.333取 1.33)
func FormatAmount(amount float64) float64 {
	// 小数点后2两位
	value := math.Floor(amount*1000) / 1000
	// 小数点后1两位
	return math.Round(value*100) / 100
}

func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}
