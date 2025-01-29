package nsqx

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"app/pkg/alert"
	"app/pkg/cliutil"
	"app/pkg/logger"
	"github.com/nsqio/go-nsq"
)

type option struct {
	maxInFlight int
	// 重试策略
	retryPolicy func(msg *nsq.Message) time.Duration
}

type Option func(*option)

func MaxInFlight(maxInFlight int) Option {
	return func(o *option) {
		o.maxInFlight = maxInFlight
	}
}

func RetryPolicy(retryPolicy func(msg *nsq.Message) time.Duration) Option {
	return func(o *option) {
		o.retryPolicy = retryPolicy
	}
}

type Payload[T any] struct {
	TraceId string `json:"trace_id"`
	Body    T      `json:"body"`
}

type plainPayload struct {
	TraceId string          `json:"trace_id"`
	Body    json.RawMessage `json:"body"`
}

func NewConsumer(addr string, topic, channel string, handle func(ctx1 context.Context, message []byte) error, ops ...Option) error {
	config := nsq.NewConfig()
	o := &option{}
	for _, op := range ops {
		op(o)
	}
	config.MaxInFlight = o.maxInFlight
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return fmt.Errorf("create nsq consumer topic: %s channel: %s error: %v", topic, channel, err)
	}
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		payload := &plainPayload{}
		err1 := json.Unmarshal(m.Body, payload)
		if err1 != nil {
			alert.Alert(context.Background(), "nsq unmarshal message body error: %v", []string{
				fmt.Sprintf("error: %s", err1.Error()),
				fmt.Sprintf("message: %s", string(m.Body)),
			})
			logger.Printf(context.Background(), "nsq unmarshal message body error: %v", err1)
			return err1
		}
		var traceId string
		if payload.TraceId != "" {
			traceId = logger.TraceIdAppendSpanId(payload.TraceId)
		}
		ctx1 := logger.NewLoggerContextWithTraceId(context.Background(), traceId)
		logger.Printf(ctx1, "nsq consumer topic(%s) channel(%s) message: %s", topic, channel, string(m.Body))
		// alert
		defer func() {
			if err1 != nil {
				alert.Alert(ctx1, "nsq consumer error: %v", []string{
					fmt.Sprintf("error: %s", err1.Error()),
				})
				logger.Printf(ctx1, "nsq consumer error: %v", err1)
			}
			if r := recover(); r != nil {
				var buf [1024]byte
				n := runtime.Stack(buf[:], false)
				alert.Alert(context.Background(), "nsq consumer panic: %v", []string{
					fmt.Sprintf("panic: %v", r),
					fmt.Sprintf("message: %s", string(m.Body)),
					fmt.Sprintf("stack: %s", string(buf[:n])),
				})
				logger.Printf(context.Background(), "nsq consumer panic: %v", r)
			}
		}()
		logger.Printf(ctx1, "receive message: %s", string(payload.Body))
		if o.retryPolicy == nil {
			if err1 = handle(ctx1, payload.Body); err1 != nil {
				logger.Printf(ctx1, "handle message error: %v", err1)
				return err1
			}
			return nil
		}
		// 禁用自动回复
		m.DisableAutoResponse()
		delay := o.retryPolicy(m)
		if err1 = handle(ctx1, payload.Body); err1 != nil {
			logger.Printf(ctx1, "handle message error: %v", err1)
			if delay > 0 {
				logger.Printf(ctx1, "requeue message after %v", delay)
				m.RequeueWithoutBackoff(delay)
				return nil
			}
		}
		m.Finish()
		return nil
	}), o.maxInFlight)
	if err = consumer.ConnectToNSQD(addr); err != nil {
		panic(fmt.Errorf("connect to nsqd error: %w", err))
	}
	cliutil.RegisterShowdown(func() {
		consumer.Stop()
		<-consumer.StopChan
	})
	return nil
}
