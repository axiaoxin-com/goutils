// viper 读取配置的常用动作封装

package goutils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitViper 根据配置文件路径和名称初始化 viper 并监听变化
// configPath 配置文件路径
// configName 配置文件名（不带格式后缀）
// configType 配置文件格式后缀
// onConfigChangeRun 配置文件发生变化时的回调函数
func InitViper(configPath, configName, configType string, onConfigChangeRun func(fsnotify.Event)) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.SetDefault("viper.inited", true)
	viper.WatchConfig()
	if onConfigChangeRun != nil {
		viper.OnConfigChange(onConfigChangeRun)
	}
	return nil
}

// IsInitedViper 返回 viper 是否已初始化
func IsInitedViper() bool {
	return viper.GetBool("viper.inited")
}
