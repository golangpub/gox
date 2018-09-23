package contents

import (
	"encoding/json"
	"github.com/gopub/types/errors"
	"reflect"
	"strings"
	"sync"
)

var mu sync.RWMutex
var contentTypeToPrototype = make(map[string]reflect.Type)
var prototypeToContentType = make(map[reflect.Type]string)

// Register bind contentType with prototype
// E.g.
//		contents.Register("image", &contents.Image{})
func Register(contentType string, prototype interface{}) error {
	contentType = strings.TrimSpace(contentType)
	if len(contentType) == 0 {
		return errors.BadRequestField("contentType")
	}

	pt := reflect.TypeOf(prototype)
	for pt.Kind() == reflect.Ptr {
		pt = pt.Elem()
	}

	if pt.Kind() != reflect.Struct {
		return errors.BadRequestField("prototype")
	}

	mu.Lock()
	defer mu.Unlock()
	if _, ok := contentTypeToPrototype[contentType]; ok {
		return errors.ConflictField("contentType")
	}

	if _, ok := prototypeToContentType[pt]; ok {
		return errors.ConflictField("prototype")
	}

	contentTypeToPrototype[contentType] = pt
	prototypeToContentType[pt] = contentType
	return nil
}

func getContentType(prototype interface{}) (string, error) {
	pt := reflect.TypeOf(prototype)
	for pt.Kind() == reflect.Ptr {
		pt = pt.Elem()
	}
	mu.RLock()
	defer mu.RUnlock()
	if contentType, ok := prototypeToContentType[pt]; ok {
		return contentType, nil
	} else {
		return "", errors.NotFoundField("prototype")
	}
}

func getProtoType(contentType string) (reflect.Type, error) {
	mu.RLock()
	defer mu.RUnlock()
	if prototype, ok := contentTypeToPrototype[contentType]; ok {
		return prototype, nil
	} else {
		return nil, errors.NotFoundField("contentType")
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
	_, err := getContentType(v)
	if err != nil {
		panic(err)
	}
	a.value = v
}

func (a *Any) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	typ, _ := m["@type"].(string)
	pt, err := getProtoType(typ)
	if err != nil {
		return errors.BadRequestField("@type")
	}

	var val = reflect.New(pt).Interface()
	err = json.Unmarshal(b, val)
	if err != nil {
		return err
	}
	a.SetValue(val)
	return nil
}

func (a *Any) MarshalJSON() ([]byte, error) {
	contentType, err := getContentType(a.value)
	if err != nil {
		return nil, errors.BadRequestField("value")
	}

	b, err := json.Marshal(a.value)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	m["@type"] = contentType
	return json.Marshal(m)
}
