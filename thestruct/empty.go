package thestruct

import (
	"reflect"
)

//
// IsEmptyValue returns true if the value is untouched
//
func IsEmptyValue(v reflect.Value) bool {

	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
