package types_test

import (
	"encoding/json"
	"fmt"
	"github.com/gopub/types"
	"testing"
	"time"
)

var _ = types.RegisterAny(&types.Image{})
var _ = types.RegisterAny(&types.Video{})
var _ = types.RegisterAny(&types.Audio{})
var _ = types.RegisterAny(&types.WebPage{})
var _ = types.RegisterAny(&types.File{})
var _ = types.RegisterAny(&types.RichText{})

func nextImage() *types.Image {
	return &types.Image{
		URL:    "https://www.image.com/" + fmt.Sprint(time.Now().Unix()),
		Width:  200,
		Height: 800,
		Format: "png",
	}
}

func nextVideo() *types.Video {
	return &types.Video{
		URL:      "http://www.video.com/" + fmt.Sprint(time.Now().Unix()),
		Format:   "rmvb",
		Duration: 1230,
		Size:     90,
		Image:    nextImage(),
	}
}

func TestID(t *testing.T) {
	var v types.ID = 10
	types.RegisterAny(types.ID(0))
	a := types.NewAny(v)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *types.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Value().(types.ID); !ok {
		t.Error("expected types.ID")
		t.FailNow()
	} else if v != v2 {
		t.Error("expected equal types.ID")
		t.FailNow()
	}
}

func TestText(t *testing.T) {
	v := "hello"
	a := types.NewAny(v)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *types.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Value().(string); !ok {
		t.Error("expected Text")
		t.FailNow()
	} else if v != v2 {
		t.Error("expected equal text value")
		t.FailNow()
	}
}

func TestImage(t *testing.T) {
	v := nextImage()
	a := types.NewAny(v)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *types.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Value().(*types.Image); !ok {
		t.Error("expected Image")
		t.FailNow()
	} else if *v != *v2 {
		t.Error("expected equal image value")
		t.FailNow()
	}
}

func TestVideo(t *testing.T) {
	v := nextVideo()
	a := types.NewAny(v)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(b))

	var a2 *types.Any
	err = json.Unmarshal(b, &a2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if v2, ok := a2.Value().(*types.Video); !ok {
		t.Error("expected Video")
		t.FailNow()
	} else if *v.Image != *v2.Image || v.URL != v2.URL || v.Size != v2.Size || v.Format != v2.Format || v.Duration != v2.Duration {
		t.Error("expected equal video value")
		t.FailNow()
	}
}

func TestArray(t *testing.T) {

	var items []*types.Any
	items = append(items, types.NewAny("hello"))
	items = append(items, types.NewAny(nextImage()))
	items = append(items, types.NewAny(nextVideo()))
	b, err := json.Marshal(items)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(string(b))

	var items2 []*types.Any
	err = json.Unmarshal(b, &items2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if items[0].Value().(string) != items2[0].Value().(string) {
		t.FailNow()
	}

	if *items[1].Value().(*types.Image) != *items2[1].Value().(*types.Image) {
		t.FailNow()
	}

	{
		v := items[2].Value().(*types.Video)
		v2 := items2[2].Value().(*types.Video)
		if *v.Image != *v2.Image || v.URL != v2.URL || v.Size != v2.Size || v.Format != v2.Format || v.Duration != v2.Duration {
			t.Error("expected equal video value")
			t.FailNow()
		}
	}
}
