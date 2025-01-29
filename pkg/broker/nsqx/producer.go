package nsqx

import (
	"context"

	"github.com/nsqio/go-nsq"
)

func NewProducer(addr string) (*nsq.Producer, error) {
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	if err = producer.Ping(); err != nil {
		return nil, err
	}
	return producer, nil
}

type Publisher interface {
	Publish(ctx context.Context, topic string, body []byte) error
	MultiPublish(ctx context.Context, topic string, body [][]byte) error
}
