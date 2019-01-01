package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
	val interface{}
}

func NewAny(v interface{}) *Any {
	a := &Any{}
	a.SetVal(v)
	return a
}

func (a *Any) Val() interface{} {
	return a.val
}

func (a *Any) SetVal(v interface{}) {
	a.val = v
}

func (a *Any) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	typ, _ := m["@type"].(string)
	pt, found := getProtoType(typ)
	if !found {
		a.val, _ = m["@value"]
		if a.val == nil {
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
	a.SetVal(ptrVal.Elem().Interface())
	return nil
}

func (a *Any) MarshalJSON() ([]byte, error) {
	name := GetAnyTypeName(a.val)

	var m = make(map[string]interface{})

	t := reflect.TypeOf(a.val)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct || t.Kind() == reflect.Map {

		b, err := json.Marshal(a.val)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}
	} else {
		m["@value"] = a.val
	}

	m["@type"] = name
	return json.Marshal(m)
}

func (a *Any) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	if b, ok := src.([]byte); ok {
		return json.Unmarshal(b, a)
	} else {
		return errors.New(fmt.Sprintf("failed to parse %v to types.Any", src))
	}
}

func (a *Any) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}

type AnyList struct {
	list []*Any
}

func NewAnyList(items ...*Any) *AnyList {
	return &AnyList{
		list: items,
	}
}

func (a *AnyList) Size() int {
	return len(a.list)
}

func (a *AnyList) Get(index int) *Any {
	return a.list[index]
}

func (a *AnyList) Remove(index int) {
	a.list = append(a.list[0:index], a.list[index+1:]...)
}

func (a *AnyList) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	if b, ok := src.([]byte); ok {
		return json.Unmarshal(b, a.list)
	} else {
		return errors.New(fmt.Sprintf("failed to parse %v to types.AnyList", src))
	}
}

func (a *AnyList) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a.list)
}
