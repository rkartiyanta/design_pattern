package structural

import (
	"reflect"
)

func Adapter(param interface{}) interface{} {
	switch v := reflect.ValueOf(param); v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	default:
		return nil
	}
}
