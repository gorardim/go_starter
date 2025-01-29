package cache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

const MaxResetPasswordTimes = 5

type UserResetPasswordLimiter struct {
	Redis *redis.Client
}

func (l *UserResetPasswordLimiter) Key(userId int) string {
	today := time.Now().Format("20060102")
	return fmt.Sprintf("user_reset_password_limiter:%d:%s", userId, today)
}

// HasChoice return true if user has choice to reset password
func (l *UserResetPasswordLimiter) HasChoice(ctx context.Context, userId int) (bool, error) {
	// 1. get key
	key := l.Key(userId)
	// 2. get value
	value, err := l.Redis.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return true, nil
		}
		return false, err
	}

	// max 3 times
	times, err := strconv.Atoi(value)
	if err != nil {
		return false, err
	}
	return times < MaxResetPasswordTimes, nil
}

// Incr user reset password times
func (l *UserResetPasswordLimiter) Incr(ctx context.Context, userId int) error {
	// get key
	key := l.Key(userId)
	// key exists
	exists, err := l.Redis.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		// 1. set
		if err = l.Redis.SetNX(ctx, key, 0, time.Hour*24).Err(); err != nil {
			return err
		}
	}
	// 2. incr
	result, err := l.Redis.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	if result > MaxResetPasswordTimes {
		return fmt.Errorf("today reset password times is used up")
	}
	return nil
}
