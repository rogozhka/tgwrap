package thestruct

import "reflect"

//
// Type is used to unwrap pointer types to their original names
//
func Type(reflectType reflect.Type) reflect.Type {
	for reflect.Slice == reflectType.Kind() || reflect.Ptr == reflectType.Kind() {
		reflectType = reflectType.Elem()
	}
	return reflectType
}
