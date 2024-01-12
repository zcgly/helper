// Version: "2023.11.29"

package helper

import (
	"reflect"
	"slices"
)

func StructToMap(item any, fields []string, otherFields ...string) map[string]any {
	fields = append(fields, otherFields...)
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	ret := make(map[string]any, len(fields))
	for i := 0; i < v.NumField(); i++ {
		ft := t.Field(i)
		key := ft.Name
		jsonTag := ft.Tag.Get("json")
		if slices.Contains(fields, jsonTag) || slices.Contains(fields, key) {
			if jsonTag != "" {
				key = jsonTag
			}
			value := v.Field(i).Interface()
			ret[key] = value
		}
	}
	return ret
}

func StructToMapStrict(item any, fields []string, otherFields ...string) map[string]any {
	ret := StructToMap(item, fields, otherFields...)
	// 检查字段数量
	if len(fields)+len(otherFields) != len(ret) {
		panic("字段名称错误")
	}
	return ret
}

func StructToMapExcept(item any, fields ...string) map[string]any {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	ret := make(map[string]any, len(fields))
	for i := 0; i < v.NumField(); i++ {
		ft := t.Field(i)
		key := ft.Name
		if slices.Contains(fields, key) {
			continue
		}
		if jsonTag := ft.Tag.Get("json"); jsonTag != "" {
			key = jsonTag
		}
		value := v.Field(i).Interface()
		ret[key] = value
	}
	return ret
}

func StructsToMaps[T any](items []T, fields []string, otherFields ...string) []map[string]any {
	ret := make([]map[string]any, len(items))
	for i, item := range items {
		ret[i] = StructToMap(item, fields, otherFields...)
	}
	return ret
}

func StructsToMapsExcept[T any](items []T, fields ...string) []map[string]any {
	ret := make([]map[string]any, len(items))
	for i, item := range items {
		ret[i] = StructToMapExcept(item, fields...)
	}
	return ret
}
