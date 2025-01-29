package email

import (
	"context"
	"fmt"
	"strings"
	"time"

	"app/pkg/randx"
	"app/services/internal/pkg/aws_ses"

	"github.com/go-redis/redis/v8"
)

type Scene string

const (
	ForgotPayPassword Scene = "FORGOT_PAY_PASSWORD"
	Register                = "REGISTER"
	Login                   = "LOGIN"
	ForgotPassword          = "FORGOT_PASSWORD"
	WithDrawApply           = "WITHDRAW_APPLY"
)

type Email struct {
	AwsSes *aws_ses.Client
	Redis  *redis.Client
}

type sendOption struct {
	subject string
	expire  time.Duration
}

type SendOption func(*sendOption)

func WithSubject(subject string) SendOption {
	return func(o *sendOption) {
		o.subject = subject
	}
}

func WithExpire(expire time.Duration) SendOption {
	return func(o *sendOption) {
		o.expire = expire
	}
}

func (e *Email) Send(ctx context.Context, email string, scene Scene, ops ...sendOption) error {
	o := &sendOption{
		subject: "inTrip: Verify Your inTrip Account",
		expire:  time.Minute * 10,
	}
	key := e.getKey(string(scene), email)
	// check ttl
	ttl, err := e.Redis.TTL(ctx, key).Result()
	if err != nil {
		return err
	}
	if ttl > 0 {
		return fmt.Errorf("email is has been sent, please check your email")
	}

	code := randx.Digit(6)
	if err := e.AwsSes.SendEmail(&aws_ses.Email{
		To:      []string{email},
		Subject: o.subject,
		Body:    strings.ReplaceAll(aws_ses.VerifyEmailTemplate, "@@@@@@", code),
	}); err != nil {
		return err
	}
	// set redis
	if err = e.Redis.Set(ctx, key, code, o.expire).Err(); err != nil {
		return err
	}
	return nil
}

func (e *Email) Verify(ctx context.Context, email string, scene Scene, code string, clear bool) error {
	// check code
	key := e.getKey(string(scene), email)
	v, err := e.Redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("email verification code is expired or not exists, please resend email")
		}
		return err
	}
	if v != code {
		return fmt.Errorf("email verification code is error, please check your email")
	}
	// clear redis
	if clear {
		if err = e.Redis.Del(ctx, key).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (e *Email) getKey(scene string, email string) string {
	return fmt.Sprintf("verify_email:%s:%s", scene, email)
}
