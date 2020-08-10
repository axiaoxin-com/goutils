package goutils

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInitViper(t *testing.T) {
	if err := InitViper(".", "viper_test", "json"); err != nil {
		t.Error("init viper return error:", err)
	}
	if id := viper.GetString("id"); id != "0001" {
		t.Error("viper should get id=0001, now id:", id)
	}
}

func TestInitViperNXFile(t *testing.T) {
	err := InitViper(".", "not_exists", "json")
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		t.Error("init viper should return ConfigFileNotFoundError when conf file not exists")
	}
}
