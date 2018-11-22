package types

type Address struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Street   string `json:"street"`
	Building string `json:"building"`
	Room     string `json:"room"`
}
