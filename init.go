package types

func init() {
	RegisterAny(&Image{})
	RegisterAny(&Audio{})
	RegisterAny(&Video{})
	RegisterAny(&File{})
	RegisterAny(&WebPage{})
	RegisterAny(&RichText{})
}
