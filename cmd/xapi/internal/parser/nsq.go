package parser

import (
	"go/ast"
	"regexp"
	"strings"

	"app/cmd/xapi/internal/utils"
)

//x:nsq order.notify 订单通知
var nsqReg = regexp.MustCompile(`^//x:nsq\s+(\S+)\s+(\S+)\s*(.*)$`)

func parseNsqTopic(text string) string {
	subMatch := nsqReg.FindStringSubmatch(text)
	if subMatch == nil {
		return ""
	}
	return subMatch[1]
}

type XChannel struct {
	Channel string
	Comment string
}

// parseXChannel parse x:channel
func parseXChannel(s string) *XChannel {
	scanner := utils.NewWordScanner(s)
	v := &XChannel{}
	if scanner.NextWord() != "//x:channel" {
		return nil
	}
	v.Channel = scanner.NextWord()
	v.Comment = scanner.Rest()
	return v
}

func nsqParse(ctx *Context, v *ast.InterfaceType) bool {
	c := &NsqConsumer{}
	c.Name = utils.GetIdentName(ctx.stack.Top(1))
	if !strings.HasSuffix(c.Name, "Consumer") {
		return false
	}
	for _, comment := range utils.GetNodeComment(ctx.stack.Top(2)) {
		if x := parseXNsq(comment); x != nil {
			c.Topic = x.Topic
			break
		}
	}
	if c.Topic == "" {
		return false
	}
	for _, field := range v.Methods.List {
		channel := &Channel{}
		funcType := field.Type.(*ast.FuncType)
		paramList := funcType.Params.List
		// check param list length
		if len(paramList) != 2 || !utils.IsContextExpr(paramList[0].Type) {
			continue
		}
		channel.Name = field.Names[0].Name
		for _, comment := range utils.GetNodeComment(field) {
			if x := parseXChannel(comment); x != nil {
				channel.Channel = x.Channel
				break
			}
		}
		if channel.Channel == "" {
			continue
		}
		channel.Request = paramList[1].Type.(*ast.StarExpr).X.(*ast.Ident).Name
		c.Request = channel.Request
		// check results length
		if len(funcType.Results.List) != 1 {
			continue
		}
		c.Channels = append(c.Channels, channel)
	}
	if len(c.Channels) > 0 {
		ctx.document.Nsq = append(ctx.document.Nsq, &Nsq{
			FileName: ctx.filename,
			Package:  ctx.pkg.Name,
			Consumers: []*NsqConsumer{
				c,
			},
		})
	}
	return true
}

type XTopic struct {
	Topic   string
	Comment string
}

func parseXNsq(s string) *XTopic {
	scanner := utils.NewWordScanner(s)
	v := &XTopic{}
	if scanner.NextWord() != "//x:nsq" {
		return nil
	}
	v.Topic = scanner.NextWord()
	v.Comment = scanner.Rest()
	return v
}
