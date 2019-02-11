package MockData

import (
	"reflect"
)

func AddTestDataToStruct(value interface{}) interface{}{

	ps := reflect.ValueOf(&value)
	s := ps.Elem()

	for i := 0; i < s.NumField(); i++ {

		//tag := val.Type().Field(i).Tag
		//fmt.Println(tag)
		//setValue := getFieldValueByTag(string(tag))
		//fmt.Println(setValue)

		switch s.Field(i).Kind() {

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if s.Field(i).CanSet() {
				s.Field(i).SetInt(345)
			}
		case reflect.String:
			if s.Field(i).CanSet() {
				s.Field(i).SetString("Bobby")
			}

		case reflect.Float32, reflect.Float64:

		}

	}

	return s

}

func getFieldValueByTag(tag string) interface{} {
	return "Hello"
}