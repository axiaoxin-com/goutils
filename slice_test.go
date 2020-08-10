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
