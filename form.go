package gox

type FormItem struct {
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Options     []string `json:"options"`
	Values      []string `json:"values"`
	Optional    bool     `json:"optional"`
	DisplayName string   `json:"display_name"`
	Description string   `json:"description"`
}

type Form struct {
	Items []*FormItem
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
