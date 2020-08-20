// viper 读取配置的常用动作封装

package goutils

import (
	"github.com/axiaoxin-com/logging"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitViper 根据配置文件路径和名称初始化 viper 并监听变化
// configPath 配置文件路径
// configName 配置文件名（不带格式后缀）
// configType 配置文件格式后缀
func InitViper(configPath, configName, configType string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logging.Warns(nil, "viper config file changed:", e.Name)
	})
	return nil
}
