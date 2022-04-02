package helpers

import (
	"reflect"
	"strconv"
)

type Decoder struct {
	Interface interface{}
	GetValue  func(string) string
}

func (d *Decoder) Decode(tag string) {
	for i := 0; i < reflect.ValueOf(d.Interface).Elem().NumField(); i++ {

		v := reflect.TypeOf(d.Interface).Elem().Field(i).Tag.Get(tag)

		obj := reflect.Indirect(reflect.ValueOf(d.Interface))

		switch obj.Field(i).Kind() {
		case reflect.String:
			obj.Field(i).SetString(d.GetValue(v))
		case reflect.Int:
			ei, err := strconv.Atoi(d.GetValue(v))
			if err != nil {
				continue
			}
			obj.Field(i).SetInt(int64(ei))
		}
	}
}
