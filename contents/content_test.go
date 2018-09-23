package contents_test

import (
	"encoding/json"
	"fmt"
	"github.com/gopub/types/contents"
	"testing"
	"time"
)

func nextText() *contents.Text {
	return &contents.Text{Text: "hello" + fmt.Sprint(time.Now().Unix())}
}

func nextImage() *contents.Image {
	return &contents.Image{
		URL:    "https://www.image.com/" + fmt.Sprint(time.Now().Unix()),
		Width:  200,
		Height: 800,
		Format: "png",
	}
}

func nextVideo() *contents.Video {
	return &contents.Video{
		URL:      "http://www.video.com/" + fmt.Sprint(time.Now().Unix()),
		Format:   "rmvb",
		Duration: 1230,
		Size:     90,
		Image:    nextImage(),
	}
}

func TestText(t *testing.T) {
	contents.Register("text", &contents.Text{})
	v := nextText()
	a := &contents.Any{Value: v}
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *contents.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Value.(*contents.Text); !ok {
		t.Error("expected Text")
		t.FailNow()
	} else if *v != *v2 {
		t.Error("expected equal text value")
		t.FailNow()
	}
}

func TestImage(t *testing.T) {
	contents.Register("image", &contents.Image{})
	v := nextImage()
	a := &contents.Any{Value: v}
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *contents.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Value.(*contents.Image); !ok {
		t.Error("expected Image")
		t.FailNow()
	} else if *v != *v2 {
		t.Error("expected equal image value")
		t.FailNow()
	}
}

func TestVideo(t *testing.T) {
	contents.Register("video", &contents.Video{})
	v := nextVideo()
	a := &contents.Any{Value: v}
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *contents.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Value.(*contents.Video); !ok {
		t.Error("expected Video")
		t.FailNow()
	} else if *v.Image != *v2.Image || v.URL != v2.URL || v.Size != v2.Size || v.Format != v2.Format || v.Duration != v2.Duration {
		t.Error("expected equal video value")
		t.FailNow()
	}
}

func TestArray(t *testing.T) {
	var items []*contents.Any
	items = append(items, &contents.Any{Value: nextText()})
	items = append(items, &contents.Any{Value: nextImage()})
	items = append(items, &contents.Any{Value: nextVideo()})
	b, err := json.Marshal(items)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(string(b))

	var items2 []*contents.Any
	err = json.Unmarshal(b, &items2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if *items[0].Value.(*contents.Text) != *items2[0].Value.(*contents.Text) {
		t.FailNow()
	}

	if *items[1].Value.(*contents.Image) != *items2[1].Value.(*contents.Image) {
		t.FailNow()
	}

	{
		v := items[2].Value.(*contents.Video)
		v2 := items2[2].Value.(*contents.Video)
		if *v.Image != *v2.Image || v.URL != v2.URL || v.Size != v2.Size || v.Format != v2.Format || v.Duration != v2.Duration {
			t.Error("expected equal video value")
			t.FailNow()
		}
	}
}
