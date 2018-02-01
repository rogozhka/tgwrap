package thestruct

import "reflect"

//
// FieldsFilterOpt represents filtering variants for Fields()
//
type FieldsFilterOpt uint

const (

	//
	// FieldsFilterDefault iterate all the fields
	// all the structures
	//
	FieldsFilterDefault = 0

	//
	// EnterEmbedsTaggedOnly by default instructs Fields()
	// to enter all the embedded(Anonymous) fields
	// and iterate their fields.
	//
	// Set EnterEmbedsTaggedOnly to true to deny entering
	// embedded fields w/o struct tag
	//
	// This rule is applied before TaggedFieldsOnly
	//
	EnterEmbedsTaggedOnly FieldsFilterOpt = 2

	// TaggedFieldsOnly by default instructs Fields()
	// to list all the fields of the struct
	//
	// Set TaggedFieldsOnly to true to skip untagged
	// struct fields
	//
	TaggedFieldsOnly FieldsFilterOpt = 1
)

//
// Fields method is used to list all the fields
// of the reflect.Type including embedded structures
//
// filterOpt: (0 by default) listing options
//
func Fields(reflectType reflect.Type, filterOpt FieldsFilterOpt) []reflect.StructField {
	var arr []reflect.StructField

	if reflectType = Type(reflectType); reflectType.Kind() == reflect.Struct {
		n := reflectType.NumField()

		for i := 0; i < n; i++ {
			v := reflectType.Field(i)

			// skip untagged fields && field is not tagged

			// is embedded
			if v.Anonymous {

				if filterOpt&EnterEmbedsTaggedOnly != 0 {
					continue
				}
				arr = append(arr, Fields(v.Type, 0)...)

			} else if filterOpt&TaggedFieldsOnly != 0 && len(v.Tag) == 0 {
				// not embedded & not tagged
				continue
			} else {
				// is not embedded & not skipped by tagged state
				arr = append(arr, v)
			}
		}
	}

	return arr
}
