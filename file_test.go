package goutils

import (
	"os"
	"testing"
)

func TestCopyFile(t *testing.T) {
	srcFilename := "viper_test.json"
	dstFilename := "viper_test.json.bak"
	if err := CopyFile(srcFilename, dstFilename); err != nil {
		t.Fatal("copy file err:", err)
	}
	defer os.Remove(dstFilename)
	if _, err := os.Stat(dstFilename); os.IsNotExist(err) {
		t.Fatal("copy file is not exists")
	}
}
