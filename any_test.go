package gox_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gopub/gox"
	"github.com/stretchr/testify/require"
)

func nextImage() *gox.Image {
	return &gox.Image{
		Link:   "https://www.image.com/" + fmt.Sprint(time.Now().Unix()),
		Width:  rand.Int(),
		Height: rand.Int(),
		Format: "png",
	}
}

func nextVideo() *gox.Video {
	return &gox.Video{
		Link:   "http://www.video.com/" + fmt.Sprint(time.Now().Unix()),
		Format: "rmvb",
		Length: rand.Int(),
		Size:   rand.Int(),
		Image:  nextImage(),
	}
}

func TestAny(t *testing.T) {
	t.Run("AliasType", func(t *testing.T) {
		gox.RegisterAny(gox.ID(0))

		v := gox.NewAny(gox.ID(10))
		jsonBytes, err := json.Marshal(v)
		require.NoError(t, err)

		var vv *gox.Any
		err = json.Unmarshal(jsonBytes, &vv)
		require.NoError(t, err)

		require.Equal(t, v.Val(), vv.Val())
	})

	t.Run("String", func(t *testing.T) {
		v := gox.NewAny("hello")
		jsonBytes, err := json.Marshal(v)
		require.NoError(t, err)

		var vv *gox.Any
		err = json.Unmarshal(jsonBytes, &vv)
		require.NoError(t, err)
		require.Equal(t, v.Val(), vv.Val())
	})

	t.Run("Struct", func(t *testing.T) {
		v := gox.NewAny(nextVideo())
		jsonBytes, err := json.Marshal(v)
		require.NoError(t, err)
		var vv *gox.Any
		err = json.Unmarshal(jsonBytes, &vv)
		require.NoError(t, err)
		require.Empty(t, gox.DiffJSON(v, vv))
	})

	t.Run("Array", func(t *testing.T) {
		var l []*gox.Any
		l = append(l, gox.NewAny("hello"))
		l = append(l, gox.NewAny(nextImage()))
		l = append(l, gox.NewAny(nextVideo()))
		jsonBytes, err := json.Marshal(l)
		require.NoError(t, err)

		var ll []*gox.Any
		err = json.Unmarshal(jsonBytes, &ll)
		require.NoError(t, err)
		require.Empty(t, gox.DiffJSON(l, ll))
	})
}
