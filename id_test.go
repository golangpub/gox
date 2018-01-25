package types

import (
	"math"
	"testing"
)

func TestID(t *testing.T) {

	t.Log(NewFastID(), NewFastID().ShortString())
	t.Log(NewID(), NewID().ShortString())
	t.Log(NewSlowID(), NewSlowID().ShortString())

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
