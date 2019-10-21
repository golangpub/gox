package gox

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type FormItem struct {
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Options     []string `json:"options"`
	Optional    bool     `json:"optional"`
	DisplayName string   `json:"display_name"`
	Description string   `json:"description"`
}

func NewFormItem() *FormItem {
	return &FormItem{}
}

func (f *FormItem) Clone() *FormItem {
	p := new(FormItem)
	*p = *f
	return p
}

type Form struct {
	Items []*FormItem `json:"items"`
}

func NewForm() *Form {
	return &Form{}
}

func (f *Form) Add(i *FormItem) {
	f.Items = append(f.Items, i)
}

func (f *Form) Remove(idx int) {
	f.Items = append(f.Items[0:idx], f.Items[idx+1:]...)
}

func (f *Form) Size() int {
	return len(f.Items)
}

func (f *Form) Get(idx int) *FormItem {
	return f.Items[idx]
}

func (f *Form) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		return json.Unmarshal([]byte(s), &f.Items)
	} else if b, ok := src.([]byte); ok {
		return json.Unmarshal(b, &f.Items)
	} else {
		return fmt.Errorf("invalid type:%v", reflect.TypeOf(src))
	}
}

func (f *Form) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}
	return json.Marshal(f.Items)
}
