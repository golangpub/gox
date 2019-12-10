package mobile

import (
	"encoding/json"

	"github.com/gopub/gox"
)

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

type AnyList struct {
	list []*Any
}

func NewAnyList() *AnyList {
	return new(AnyList)
}

func (a *AnyList) Size() int {
	if a == nil {
		return 0
	}
	return len(a.list)
}

func (a *AnyList) Get(index int) *Any {
	if a == nil {
		return nil
	}
	return a.list[index]
}

func (a *AnyList) Append(v *Any) {
	a.list = append(a.list, v)
}

func (a *AnyList) Prepend(v *Any) {
	a.list = append([]*Any{v}, a.list...)
}

func (a *AnyList) Insert(i int, v *Any) {
	if len(a.list) <= i {
		a.list = append(a.list, v)
	} else {
		l := a.list[i:]
		l = append([]*Any{v}, l...)
		a.list = append(a.list[0:i], l...)
	}
}

func (a *AnyList) RemoveAt(index int) {
	a.list = append(a.list[0:index], a.list[index+1:]...)
}

func (a *AnyList) Remove(v *Any) {
	i := a.IndexOf(v)
	if i >= 0 {
		a.RemoveAt(i)
	}
}

func (a *AnyList) IndexOf(v *Any) int {
	for i, m := range a.list {
		if m == v {
			return i
		}
	}
	return -1
}

func (a *AnyList) FirstImage() *Image {
	for _, m := range a.list {
		if img := m.Image(); img != nil {
			return img
		}
	}
	return nil
}

func (a *AnyList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &a.list)
}

func (a *AnyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.list)
}
