package goutils

import (
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func TestInitViper(t *testing.T) {
	defer viper.Reset()
	onConfigChangeRun := func(e fsnotify.Event) {
		t.Log("run when config file is changed")
	}
	if err := InitViper(".", "viper.webapp", "toml", onConfigChangeRun); err != nil {
		t.Error("init viper return error:", err)
	}
	if viper.GetString("env") == "" {
		t.Error("viper should get env value")
	}
	if !IsInitedViper() {
		t.Error("viper inited should return true")
	}
}

func TestInitViperNXFile(t *testing.T) {
	defer viper.Reset()
	err := InitViper(".", "not_exists", "json", nil)
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		t.Error("init viper should return ConfigFileNotFoundError when conf file not exists")
	}
	if IsInitedViper() {
		t.Error("viper inited should return false")
	}
}

func TestInitWebAppViperConfig(t *testing.T) {
	defer viper.Reset()
	InitWebAppViperConfig()
	if viper.GetString("server.addr") == "" {
		t.Error("webapp config init failed")
	}
}
