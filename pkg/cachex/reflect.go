package cachex

import "reflect"

// a key is from struct type

func getKeyFromType(t reflect.Type) string {
	return t.PkgPath() + "." + t.Name()
}
