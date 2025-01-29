package cache

import (
	"context"
	"fmt"
	"time"

	"app/component/locks"
)

type OneKeyReceiveIncomeLock struct {
	RedisLock *locks.RedisLock
}

func (l *OneKeyReceiveIncomeLock) Lock(ctx context.Context, userId int) error {
	_, err := l.RedisLock.Lock(ctx, l.getKey(userId), time.Minute)
	return err
}

func (l *OneKeyReceiveIncomeLock) Unlock(userId int) error {
	return l.RedisLock.Unlock(context.Background(), l.getKey(userId))
}

func (l *OneKeyReceiveIncomeLock) getKey(userId int) string {
	return fmt.Sprintf("onekey_receive_income_lock:%d", userId)
}
