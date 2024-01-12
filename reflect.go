package helper

import "reflect"

func GetStructFields(item any) []reflect.StructField {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	var ret []reflect.StructField
	for i := 0; i < v.NumField(); i++ {
		ret = append(ret, t.Field(i))
	}
	return ret
}

func GetStructFieldNames(item any) []string {
	fields := GetStructFields(item)
	ret := make([]string, 0, len(fields))
	for _, field := range fields {
		ret = append(ret, field.Name)
	}
	return ret
}

func GetStructJsons(item any) []string {
	fields := GetStructFields(item)
	ret := make([]string, 0, len(fields))
	for _, field := range fields {
		json := field.Tag.Get("json")
		if json == "-" {
			continue
		}
		ret = append(ret, json)
	}
	return ret
}
