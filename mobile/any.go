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
	List []*Any
}

func NewAnyList(list []*Any) *AnyList {
	return new(AnyList)
}

func (a *AnyList) Size() int {
	if a == nil {
		return 0
	}
	return len(a.List)
}

func (a *AnyList) Get(index int) *Any {
	if a == nil {
		return nil
	}
	return a.List[index]
}

func (a *AnyList) Append(v *Any) {
	a.List = append(a.List, v)
}

func (a *AnyList) Prepend(v *Any) {
	a.List = append([]*Any{v}, a.List...)
}

func (a *AnyList) Insert(i int, v *Any) {
	if len(a.List) <= i {
		a.List = append(a.List, v)
	} else {
		l := a.List[i:]
		l = append([]*Any{v}, l...)
		a.List = append(a.List[0:i], l...)
	}
}

func (a *AnyList) RemoveAt(index int) {
	a.List = append(a.List[0:index], a.List[index+1:]...)
}

func (a *AnyList) Remove(v *Any) {
	i := a.IndexOf(v)
	if i >= 0 {
		a.RemoveAt(i)
	}
}

func (a *AnyList) IndexOf(v *Any) int {
	for i, m := range a.List {
		if m == v {
			return i
		}
	}
	return -1
}

func (a *AnyList) FirstImage() *Image {
	for _, m := range a.List {
		if img := m.Image(); img != nil {
			return img
		}
	}
	return nil
}

func (a *AnyList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &a.List)
}

func (a *AnyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.List)
}
