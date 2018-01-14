package thestruct

import (
	"strings"
)

func explodeStructTag(tag string) ([]*Tag, error) {

	tag = strings.TrimSpace(tag)

	var res []*Tag

	// json:"name,string" xml:"name" gorm:"primary_key" bigquery:",nullable"

	arr := strings.Split(tag, " ")

	for i := 0; i < len(arr); i++ {
		t := strings.TrimSpace(arr[i])

		key, val := getStructNameValue(t)

		//val := getStructValue(strings.TrimSpace(arr[i+1]))

		arr2 := strings.Split(val, ",")

		if len(arr2) < 1 {
			continue
		}
		name := arr2[0]

		newItem := Tag{
			Section: key,
			Value:   name,
			mapOpt:  make(map[string]bool),
		}

		for j := 1; j < len(arr2); j++ {
			optVal := strings.TrimSpace(arr2[j])
			newItem.Options = append(newItem.Options, optVal)
			newItem.mapOpt[optVal] = true
		}

		res = append(res, &newItem)
	}

	return res, nil
}

func getStructValue(val string) string {

	offset := strings.Index(val, " ")

	if -1 == offset {
		return val
	}

	return val[:offset]
}

func getStructNameValue(s string) (string, string) {

	arr := strings.Split(s, ":")

	if len(arr) < 2 {
		return "", ""
	}

	name := arr[0]
	rawVal := arr[1]
	runes := []rune(rawVal)

	if len(runes) < 3 || runes[0] != '"' || runes[len(runes)-1] != '"' {
		return name, ""
	}

	return name, rawVal[1 : len(rawVal)-1]
}
