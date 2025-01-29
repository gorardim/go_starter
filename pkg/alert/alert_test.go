package alert

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestAlert(t *testing.T) {
	Alert(context.Background(), "test", []string{
		fmt.Sprintf("错误信息：%s", "test"),
	})

	time.Sleep(time.Second * 5)
}
