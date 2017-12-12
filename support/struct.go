package support

import (
	"reflect"
	"strconv"
)

const paramTagName = "param"

func Map(s interface{}) (map[string]string) {

	v := reflect.ValueOf(s)

	m := make(map[string]string)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get(paramTagName)

		var value string

		switch v.Field(i).Interface().(type) {
		case string:
			value = v.Field(i).String()
		case int:
			value = strconv.Itoa(int(v.Field(i).Int()))
		case float64:
			value = strconv.FormatFloat(v.Field(i).Float(), 'f', -1, 64)
		case bool:
			if v.Field(i).Bool() {
				value = "1"
			} else {
				value = "0"
			}
		default:
			value = v.Field(i).String()
		}

		m[tag] = value
	}

	return m
}
