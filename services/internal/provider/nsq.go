package provider

import (
	"app/pkg/broker/nsqx"
	"app/services/internal/config"

	"github.com/nsqio/go-nsq"
)

func NewNsqProducer(conf *config.Config) *nsq.Producer {
	producer, err := nsqx.NewProducer(conf.Nsq.Addr)
	if err != nil {
		panic(err)
	}
	return producer
}
