// viper 读取配置的常用动作封装

package goutils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// OnVipConfigChangeRun vip config文件变化时的回调函数类型
type OnVipConfigChangeRun func(fsnotify.Event)

// InitViper 根据配置文件路径和名称初始化 viper 并监听变化
// configPath 配置文件路径
// configName 配置文件名（不带格式后缀）
// configType 配置文件格式后缀
// onConfigChangeRun 配置文件发生变化时的回调函数
func InitViper(configPath, configName, configType string, run OnVipConfigChangeRun) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	if run != nil {
		viper.OnConfigChange(run)
	}
	return nil
}
