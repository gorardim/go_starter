package cachex

import (
	"reflect"
	"testing"
)

func Test_getKeyFromType(t *testing.T) {
	t.Log(getKeyFromType(reflect.TypeOf(Flight[int]{})))
}
