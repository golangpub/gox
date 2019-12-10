package mobile

import "github.com/gopub/gox"

type Video = gox.Video

func NewVideo() *Video {
	return new(Video)
}

type Audio = gox.Audio

func NewAudio() *Audio {
	return new(Audio)
}

type File = gox.File

func NewFile() *File {
	return new(File)
}

type WebPage = gox.WebPage

func NewWebPage() *WebPage {
	return new(WebPage)
}

type Any = gox.Any

func NewAnyObj() *Any {
	return new(Any)
}

type AnyList = gox.AnyList

func NewAnyList() *AnyList {
	return new(AnyList)
}
