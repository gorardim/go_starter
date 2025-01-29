package cache

import (
	"context"
	"fmt"
	"time"

	"app/component/counter"
	"app/pkg/alert"
)

type PayPasswordLimiter struct {
	Counter *counter.Counter
}

func (l *PayPasswordLimiter) Check(ctx context.Context, userId int) bool {
	get, err := l.Counter.Get(ctx, fmt.Sprintf("pay_password_limiter:%d", userId))
	if err != nil {
		alert.Alert(ctx, "get pay password limiter error", []string{
			fmt.Sprintf("user_id: %d", userId),
			fmt.Sprintf("error: %s", err),
		})
		return false
	}
	return get >= 5
}

func (l *PayPasswordLimiter) Incr(ctx context.Context, userId int) {
	_, err := l.Counter.Incr(ctx, fmt.Sprintf("pay_password_limiter:%d", userId), time.Minute*10)
	if err != nil {
		alert.Alert(ctx, "incr pay password limiter error", []string{
			fmt.Sprintf("user_id: %d", userId),
			fmt.Sprintf("error: %s", err),
		})
	}
}
