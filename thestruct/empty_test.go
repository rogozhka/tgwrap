package thestruct

import (
	"reflect"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

type IsEmptyTestCase struct {
	Inbound  interface{}
	Expected bool
}

func TestIsEmptyValue(t *testing.T) {

	var ptr *Tag
	var ptrUsed *Tag = &Tag{}
	var arrEmpty []int
	var arrNonEmpty []int
	arrNonEmpty = append(arrNonEmpty, 123)

	type structA struct {
		A int
		B float32
	}

	type structB struct {
		structA
		C int
	}

	stInst := structB{}
	stInst2 := structB{
		structA: structA{
			A: 0,
		},
	}
	stInst3 := structB{
		structA: structA{},
	}

	stInst4 := structB{
		structA: structA{
			B: 13.5,
		},
	}

	arr := []IsEmptyTestCase{
		{
			0,
			true,
		},
		{
			uint(0),
			true,
		},
		{
			int(1),
			false,
		},
		{
			ptr,
			true,
		},
		{
			ptrUsed,
			false,
		},
		{
			arrEmpty,
			true,
		},
		{
			arrNonEmpty,
			false,
		},
		{
			stInst,
			true,
		},
		{
			stInst2,
			true,
		},
		{
			stInst3,
			true,
		},
		{
			stInst4,
			false,
		},
	}

	for _, el := range arr {

		assert.Equal(t, el.Expected, IsEmptyValue(reflect.ValueOf(el.Inbound)),
			fmt.Sprintf("IsEmptyValue:%v", el.Inbound))
	}

}
