package config

import (
	"os"
	"reflect"
)

const tag string = "env"

func EnvEncode(s interface{}) {
	for i := 0; i < reflect.ValueOf(s).Elem().NumField(); i++ {

		v := reflect.TypeOf(s).Elem().Field(i).Tag.Get(tag)

		obj := reflect.Indirect(reflect.ValueOf(s))
		obj.Field(i).SetString(os.Getenv(v))
	}
}
