package cliutil

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"reflect"
)

var cliCmdType = reflect.TypeOf((*cli.Command)(nil))

// NewCliCommand converts a struct to a cli.Command list.
func NewCliCommand(itemStruct interface{}) []*cli.Command {
	var commands []*cli.Command
	// must ptr
	rv := reflect.ValueOf(itemStruct)
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("cli commands must ptr, got %v", rv.Kind()))
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		panic(fmt.Sprintf("cli commands must struct, got %v", rv.Kind()))
	}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if field.IsNil() {
			continue
		}
		command := field.Convert(cliCmdType).Interface().(*cli.Command)
		commands = append(commands, command)
	}
	return commands
}
