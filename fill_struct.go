package TestSweet

import (
	"reflect"
)

func FillStruct(s interface{}) {
	v := reflect.ValueOf(s).Elem()
	t := reflect.TypeOf(s).Elem()

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)
		varType := ft.Tag.Get("test")

		if varType != "" || fv.Kind() == reflect.Struct {
			switch fv.Kind() {
			case reflect.String:
				if fv.CanSet() {
					value := GetStringValue(varType)
					fv.SetString(value)
				}
				break
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				if fv.CanSet() {
					value := GetIntValue(varType)
					fv.SetInt(value)
				}
				break
			case reflect.Float32, reflect.Float64:
				if fv.CanSet() {
					value := GetFloatValue(varType)
					fv.SetFloat(value)
				}
				break
			}
		}
	}
}
