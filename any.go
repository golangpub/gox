package types

import (
	"encoding/json"
	"errors"
	"reflect"
	"sync"
)

var mu sync.RWMutex
var nameToPrototype = make(map[string]reflect.Type)

type AnyType interface {
	AnyType() string
}

// Register bind typ with prototype
// E.g.
//		contents.Register("image", &contents.Image{})
func RegisterAny(prototype interface{}) error {
	name := GetAnyTypeName(prototype)
	mu.Lock()
	defer mu.Unlock()
	if _, ok := nameToPrototype[name]; ok {
		return errors.New("conflict type name: " + name)
	}

	nameToPrototype[name] = reflect.TypeOf(prototype)
	return nil
}

func camelToSnake(s string) string {
	snake := make([]rune, 0, len(s)+1)
	flag := false
	k := 'a' - 'A'
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if !flag {
				flag = true
				if i > 0 {
					snake = append(snake, '_')
				}
			}
			snake = append(snake, c+k)
		} else {
			flag = false
			snake = append(snake, c)
		}
	}
	return string(snake)
}

func GetAnyTypeName(prototype interface{}) string {
	if a, ok := prototype.(AnyType); ok {
		return a.AnyType()
	}

	p := reflect.TypeOf(prototype)
	for p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	return camelToSnake(p.Name())
}

func getProtoType(typ string) (reflect.Type, bool) {
	mu.RLock()
	defer mu.RUnlock()
	if prototype, ok := nameToPrototype[typ]; ok {
		return prototype, true
	} else {
		return nil, false
	}
}

type Any struct {
	value interface{}
}

func NewAny(v interface{}) *Any {
	a := &Any{}
	a.SetValue(v)
	return a
}

func (a *Any) Value() interface{} {
	return a.value
}

func (a *Any) SetValue(v interface{}) {
	a.value = v
}

func (a *Any) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	typ, _ := m["@type"].(string)
	pt, found := getProtoType(typ)
	if !found {
		a.value, _ = m["@value"]
		if a.value == nil {
			return errors.New("value is empty")
		}
		return nil
	}

	if v, ok := m["@value"]; ok {
		b, _ = json.Marshal(v)
	}

	var ptrVal = reflect.New(pt)

	for val := ptrVal; val.Kind() == reflect.Ptr && val.CanSet(); val = val.Elem() {
		val.Set(reflect.New(val.Elem().Type()))
	}

	err := json.Unmarshal(b, ptrVal.Interface())
	if err != nil {
		return err
	}
	a.SetValue(ptrVal.Elem().Interface())
	return nil
}

func (a *Any) MarshalJSON() ([]byte, error) {
	name := GetAnyTypeName(a.value)

	var m = make(map[string]interface{})

	t := reflect.TypeOf(a.value)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct || t.Kind() == reflect.Map {

		b, err := json.Marshal(a.value)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}
	} else {
		m["@value"] = a.value
	}

	m["@type"] = name
	return json.Marshal(m)
}
