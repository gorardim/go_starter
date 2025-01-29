package cronx

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"app/pkg/alert"
	"app/pkg/logger"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	c *cron.Cron
}

func NewCron() *Cron {
	return &Cron{
		c: cron.New(),
	}
}

func (c *Cron) AddFunc(name string, spec string, fn func(context.Context) error) {
	log.Printf("cron add func: %s, spec: %s\n", name, spec)
	_, err := c.c.AddFunc(spec, func() {
		ctx := logger.NewLoggerContext(context.Background())
		// recover
		defer func() {
			if r := recover(); r != nil {
				var buf [1024]byte
				n := runtime.Stack(buf[:], false)
				alert.Alert(context.Background(), "cron job panic", []string{
					fmt.Sprintf("panic: %v", r),
					fmt.Sprintf("name: %s", name),
					fmt.Sprintf("spec: %s", spec),
					fmt.Sprintf("stack: %s", string(buf[:n])),
				})
			}
		}()
		if err := fn(ctx); err != nil {
			alert.Alert(ctx, "cron", []string{
				"name: " + name,
				"cron error: " + err.Error(),
			})
		}
	})

	if err != nil {
		log.Fatalf("cron %s add func error: %s", name, err.Error())
	}
}

func (c *Cron) Run() {
	log.Println("cron start")
	c.c.Run()
}
