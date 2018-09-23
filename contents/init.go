package contents

func init() {
	Register("text", &Text{})
	Register("image", &Image{})
	Register("audio", &Audio{})
	Register("video", &Video{})
	Register("file", &File{})
	Register("web_page", &WebPage{})
}
