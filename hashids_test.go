package goutils

import "testing"

func TestAll(t *testing.T) {
	h, err := NewHashids("salt", 8, "")
	if err != nil {
		t.Error(err)
	}

	var id int64 = 99999999999999999
	str, err := h.Encode(id)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%d encode to %s", id, str)

	id, err = h.Decode(str)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s decode to %d", str, id)

	h, err = NewHashids("salt", 8, "pre-")
	if err != nil {
		t.Error(err)
	}

	id = 99999999999999999
	str, err = h.Encode(id)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%d encode to %s", id, str)

	id, err = h.Decode(str)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s decode to %d", str, id)

}
