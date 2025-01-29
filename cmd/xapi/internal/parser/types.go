package parser

type Document struct {
	Services []*Service
	Schemas  map[string]*Schema
	Nsq      []*Nsq
}

func newDocument() *Document {
	return &Document{
		Services: make([]*Service, 0),
		Schemas:  make(map[string]*Schema),
	}
}

type Service struct {
	Name     string
	Comment  []string
	RpcArr   []*Rpc
	FileName string
	Package  string
}

type Rpc struct {
	Name     string
	Comment  []string
	Request  string
	Response string
	Path     string
	Method   string
}

type Schema struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name    string
	Type    string
	Comment []string
	Tags    map[string]string
}

type Nsq struct {
	Package   string
	FileName  string
	Consumers []*NsqConsumer
}

type NsqConsumer struct {
	Topic    string
	Name     string
	Request  string
	Channels []*Channel
}

type Channel struct {
	Name    string
	Request string
	Channel string
}
