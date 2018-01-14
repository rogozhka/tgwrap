package thestruct

import "reflect"

func Fields(reflectType reflect.Type) []reflect.StructField {
	var arr []reflect.StructField

	if reflectType = Type(reflectType); reflectType.Kind() == reflect.Struct {
		n := reflectType.NumField()

		for i := 0; i < n; i++ {
			v := reflectType.Field(i)
			if len(v.Tag) > 0 && v.Anonymous {
				arr = append(arr, Fields(v.Type)...)
			} else {
				arr = append(arr, v)
			}
		}
	}

	return arr
}

func Type(reflectType reflect.Type) reflect.Type {
	for reflect.Slice == reflectType.Kind() || reflect.Ptr == reflectType.Kind() {
		reflectType = reflectType.Elem()
	}
	return reflectType
}
