package alert

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"app/pkg/logger"
)

var Env = ""

type options struct {
	sync bool
}
type Option func(*options)

func WithSync() Option {
	return func(o *options) {
		o.sync = true
	}
}

func Alert(ctx context.Context, title string, args []string, ops ...Option) {
	o := &options{}
	for _, op := range ops {
		op(o)
	}
	if Env == "" {
		logger.Printf(ctx, "%s\n%s", title, strings.Join(args, "\n"))
		return
	}
	buf := &bytes.Buffer{}
	buf.WriteString(title)
	buf.WriteString(":")
	buf.WriteString("\n")
	for _, arg := range args {
		buf.WriteString("- ")
		buf.WriteString(arg)
		buf.WriteString("\n")
	}
	buf.WriteString("- ")
	buf.WriteString("报警时间: ")
	buf.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	buf.WriteString("\n")
	// trace id
	buf.WriteString("- ")
	buf.WriteString("trace id: ")
	buf.WriteString(logger.TraceIdFromLogger(ctx))
	buf.WriteString("\n")
	// env
	buf.WriteString("- ")
	buf.WriteString("env: ")
	buf.WriteString(Env)
	if o.sync {
		sendMessage(buf.String())
		return
	}
	select {
	case queue <- buf.String():
	default:
		log.Printf("alert queue is full")
	}
}

type remoteAlert struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

func sendMessage(text string) {
	v := remoteAlert{
		ChatId: "-4122894274",
		Text:   text,
	}
	body, err := json.Marshal(v)
	if err != nil {
		log.Printf("send alert error: %s", err)
		return
	}
	resp, err := http.Post("https://api.telegram.org/botBOT:ID/sendMessage", "application/json", bytes.NewReader(body))
	if err != nil {
		log.Printf("send alert error: %s", err)
		return
	}
	defer resp.Body.Close()
}
