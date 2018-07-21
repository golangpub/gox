package types

import (
	"math"
	"testing"
)

func TestID(t *testing.T) {
	for i := 0; i < 256; i++ {
		id := NextID()
		t.Logf("%d %0X %s", id, id, id.ShortString())
		//time.Sleep(time.Millisecond * 1)
	}

	var id ID = 123
	if id.ShortString() != "1z" {
		t.Log(id.ShortString())
		t.FailNow()
	}

	id = 62
	if id.ShortString() != "10" {
		t.Log(id.ShortString())
		t.FailNow()
	}

	id = math.MaxInt64
	if i, _ := ParseShortID(id.ShortString()); i != id {
		t.Log(id.ShortString(), i)
		t.FailNow()
	}
}
