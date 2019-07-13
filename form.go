package gox

type FormItem struct {
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Options     []string `json:"options"`
	Values      []string `json:"values"`
	Optional    bool     `json:"optional"`
	Description string   `json:"description"`
}
