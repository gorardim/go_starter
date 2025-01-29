package job

import (
	"app/pkg/broker/nsqx"
	"app/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

{{range $c := .Consumers}}
    {{range $ch := $c.Channels}}
    func {{trimSuffix $c.Name "Consumer"}}{{$ch.Name}}Consumer(svc {{$c.Name}}, addr string, options ...nsqx.Option)  {
        err := nsqx.NewConsumer(addr,"{{$c.Topic}}", "{{$ch.Channel}}", func(ctx context.Context, message []byte) error {
            req := &{{$ch.Request}}{}
            if err := json.Unmarshal(message, req); err != nil {
                return fmt.Errorf("unmarshal message error: %w", err)
            }
            return svc.{{$ch.Name}}(ctx,req)
        }, options...)
        if err != nil {
            panic(fmt.Errorf("create consumer error: %w", err))
        }
    }
   {{end}}

    type {{trimSuffix $c.Name "Consumer"}}Publisher interface {
        Publish(ctx context.Context, msg *{{trimSuffix $c.Name "Consumer"}}Request) error
        DeferredPublish(ctx context.Context, msg *{{trimSuffix $c.Name "Consumer"}}Request, delay time.Duration) error
        MultiPublish(ctx context.Context, msg []*{{trimSuffix $c.Name "Consumer"}}Request) error
    }

    var _ {{trimSuffix $c.Name "Consumer"}}Publisher = (*{{trimSuffix $c.Name "Consumer"}}PublisherImpl)(nil)

    type {{trimSuffix $c.Name "Consumer"}}PublisherImpl struct {
        producer *nsq.Producer
    }

    func New{{trimSuffix $c.Name "Consumer"}}Publisher(producer *nsq.Producer) {{trimSuffix $c.Name "Consumer"}}Publisher {
        return &{{trimSuffix $c.Name "Consumer"}}PublisherImpl{
            producer: producer,
        }
    }

    func (d *{{trimSuffix $c.Name "Consumer"}}PublisherImpl) Publish(ctx context.Context, msg *{{trimSuffix $c.Name "Consumer"}}Request) error {
        data, err := json.Marshal(&nsqx.Payload[{{trimSuffix $c.Name "Consumer"}}Request]{
            TraceId: logger.TraceIdFromLogger(ctx),
            Body:    *msg,
        })
        if err != nil {
            return err
        }
        return d.producer.Publish("{{$c.Topic}}", data)
    }

    func (d *{{trimSuffix $c.Name "Consumer"}}PublisherImpl) DeferredPublish(ctx context.Context, msg *{{trimSuffix $c.Name "Consumer"}}Request, delay time.Duration) error {
        data, err := json.Marshal(&nsqx.Payload[{{trimSuffix $c.Name "Consumer"}}Request]{
            TraceId: logger.TraceIdFromLogger(ctx),
            Body:    *msg,
        })
        if err != nil {
            return err
        }
        return d.producer.DeferredPublish("{{$c.Topic}}", delay, data)
    }

    func (d *{{trimSuffix $c.Name "Consumer"}}PublisherImpl) MultiPublish(ctx context.Context, msg []*{{trimSuffix $c.Name "Consumer"}}Request) error {
        data := make([][]byte, 0, len(msg))
        for _, v := range msg {
            b, err := json.Marshal(&nsqx.Payload[{{trimSuffix $c.Name "Consumer"}}Request]{
                TraceId: logger.TraceIdFromLogger(ctx),
                Body:    *v,
            })
            if err != nil {
                return err
            }
            data = append(data, b)
        }
        return d.producer.MultiPublish("{{$c.Topic}}", data)
    }
{{end}}