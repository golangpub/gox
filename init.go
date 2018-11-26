package types

func init() {
	RegisterAny(&Image{})
	RegisterAny(&Video{})
	RegisterAny(&Audio{})
	RegisterAny(&WebPage{})
	RegisterAny(&File{})
	RegisterAny(&RichText{})
}
