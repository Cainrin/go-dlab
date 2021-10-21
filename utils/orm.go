package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Cainrin/go-dlab/errors"
	"github.com/gogf/gf/frame/g"
	"reflect"
	"strings"
)

type PqSlice string

func (c PqSlice) MarshalJSON() ([]byte, error) {
	s := strings.Split(string(c), ";")
	m, err := json.Marshal(s)
	return m, err
}

func (c PqSlice) ToArray() []string {
	s := strings.Split(strings.ReplaceAll(string(c), "'", ""), ";")
	return s
}

func WhereMap(data interface{}, customWhere g.Map) g.Map {
	var whereMap = g.Map{}
	if customWhere != nil {
		for k, v := range customWhere {
			whereMap[k] = v
		}
	}
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Struct {
		panic("panic: need struct param")
	}
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldTagOrm := field.Tag.Get("orm")
		value := v.Field(i).Interface()
		// only struct tag contains orm and filed not Zero val will be
		if reflect.ValueOf(fieldTagOrm).IsZero() || reflect.ValueOf(value).IsZero() {
			continue
		}
		whereMap[fieldTagOrm] = value
	}
	return whereMap
}

func Where(data interface{}) g.Map {
	var whereMap = g.Map{}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	switch t.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		v = v.Elem()
	default:
		panic("panic: need struct  or struct pointer")
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldTagOrm := field.Tag.Get("orm")
		value := v.Field(i).Interface()
		// only struct tag contains orm and filed not Zero val will be
		if !reflect.ValueOf(fieldTagOrm).IsZero() && !reflect.ValueOf(value).IsZero() {
			whereMap[fieldTagOrm] = value
		}
	}
	return whereMap
}
func InsertMap(data interface{}) g.Map {
	var whereMap = g.Map{}
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Struct {
		panic("panic: data need struct type!")
	}
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fieldTagOrm := field.Tag.Get("orm")
		fieldTagV := field.Tag.Get("v")
		if reflect.ValueOf(value).IsZero() &&
			reflect.ValueOf(fieldTagV).String() == "required" {
			panic(fmt.Sprintf("%s is required", field.Name))
		}
		// only struct tag contains orm and filed not Zero val
		if reflect.ValueOf(fieldTagOrm).IsZero() || reflect.ValueOf(value).IsZero() {
			continue
		}
		// slice change to postgres array
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			value = SpliceToPgString(value)
		}
		whereMap[fieldTagOrm] = value
	}
	return whereMap
}

func IsObjectExist(tableName string, data interface{}) bool {
	if count, err := g.DB("default").Model(tableName).Count(data); err != nil {
		return false
	} else {
		return count != 0
	}
}

func SetWithTimeOut(key string, val interface{}, timeout int) error {
	if _, err := g.Redis().Do("SET", key, val); err != nil {
		return errors.ORMError("redis set error")
	}

	if _, err := g.Redis().Do("EXPIRE", key, timeout); err != nil {
		return errors.ORMError("redis expire error")
	}
	return nil
}
