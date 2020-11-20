package goutils

import "testing"

func TestIsEqualStringSlice(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "b", "c"}
	s3 := []string{"a", "d", "c"}
	if !IsEqualStringSlice(s1, s2) {
		t.Error("s1 == s2, should return true")
	}

	if IsEqualStringSlice(s1, s3) {
		t.Error("s1 != s3, should return false")
	}
}

func TestRemoveStringSliceItemByIndex(t *testing.T) {
	rawStrSlice := []string{"0", "1", "2", "3"}
	newStrSlice := RemoveStringSliceItemByIndex(rawStrSlice, 1)
	expectResult := []string{"0", "2", "3"}
	if !IsEqualStringSlice(newStrSlice, expectResult) {
		t.Error("RemoveStringSliceItemByIndex return wrong result:", "raw:", rawStrSlice, "new:", newStrSlice)
	}
}

func TestChunkFloat64Slice(t *testing.T) {
	raw := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := ChunkFloat64Slice(raw, 3)
	if len(result) != 3 {
		t.Error("chunk 3 error", result)
	}
	result = ChunkFloat64Slice(raw, 4)
	if len(result) != 3 {
		t.Error("chunk 4 error", result)
	}
	result = ChunkFloat64Slice(raw, 5)
	if len(result) != 2 {
		t.Error("chunk 5 error", result)
	}
}

func TestIsStrInSlice(t *testing.T) {
	r := IsStrInSlice("asd", []string{"xx", "aa", "xasd"})
	if r == true {
		t.Error("should return false")
	}
	r = IsStrInSlice("asd", []string{"xx", "aa", "asd"})
	if r == false {
		t.Error("should return true")
	}
}

func TestIsIntInSlice(t *testing.T) {
	r := IsIntInSlice(1, []int{7, 30})
	if r == true {
		t.Error("should return false")
	}
	r = IsIntInSlice(2, []int{2, 3, 4})
	if r == false {
		t.Error("should return true")
	}
}
