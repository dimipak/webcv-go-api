package config

import (
	"os"
	"reflect"
	"strconv"
)

const tag string = "env"

type Config interface {
	setValues()
}

func setValues(c ...Config) {
	for _, v := range c {
		v.setValues()
	}
}

type Encode struct {
	Struct   interface{}
	GetValue func(string) string
}

func envEncode(s interface{}) {
	for i := 0; i < reflect.ValueOf(s).Elem().NumField(); i++ {

		v := reflect.TypeOf(s).Elem().Field(i).Tag.Get(tag)

		obj := reflect.Indirect(reflect.ValueOf(s))

		switch obj.Field(i).Kind() {
		case reflect.String:
			obj.Field(i).SetString(os.Getenv(v))
		case reflect.Int:
			ei, err := strconv.Atoi(os.Getenv(v))
			if err != nil {
				continue
			}
			obj.Field(i).SetInt(int64(ei))
		}
	}
}

func (e *Encode) EnvEncode2(tag2 string) {
	for i := 0; i < reflect.ValueOf(e.Struct).Elem().NumField(); i++ {

		v := reflect.TypeOf(e.Struct).Elem().Field(i).Tag.Get(tag2)

		obj := reflect.Indirect(reflect.ValueOf(e.Struct))

		switch obj.Field(i).Kind() {
		case reflect.String:
			obj.Field(i).SetString(e.GetValue(v))
		case reflect.Int:
			ei, err := strconv.Atoi(e.GetValue(v))
			if err != nil {
				continue
			}
			obj.Field(i).SetInt(int64(ei))
		}
	}
}
