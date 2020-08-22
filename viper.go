// viper 读取配置的常用动作封装

package goutils

import (
	"flag"
	"os"

	"github.com/axiaoxin-com/logging"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
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

// InitWebAppViperConfig 按 viper webapp 配置文件初始化 viper 配置
func InitWebAppViperConfig() {
	// 加载配置文件到 viper
	workdir, err := os.Getwd()
	if err != nil {
		logging.Warn(nil, "get workdir failed:"+err.Error())
		workdir = "."
	}
	configPath := flag.String("p", workdir, "path of config file")
	configName := flag.String("c", "viper.webapp", "name of config file without format suffix)")
	configType := flag.String("t", "toml", "type of config file format")
	flag.Parse()
	if err := InitViper(*configPath, *configName, *configType, func(e fsnotify.Event) {
		logging.Warn(nil, "Config file changed:"+e.Name)
	}); err != nil {
		logging.Warn(nil, "Init viper error:"+err.Error())
	}

	// 设置配置默认值
	viper.SetDefault("env", "dev")

	viper.SetDefault("server.addr", ":4869")
	viper.SetDefault("server.mode", gin.ReleaseMode)
	viper.SetDefault("server.pprof", true)
	viper.SetDefault("server.read_timeout", 5)  // 服务器从 accept 到读取 body 的超时时间（秒）
	viper.SetDefault("server.write_timeout", 5) // 服务器从 accept 到写 response 的超时时间（秒）

	viper.SetDefault("apidocs.title", "pink-lady swagger apidocs")
	viper.SetDefault("apidocs.desc", "Using pink-lady to develop gin app on fly.")
	viper.SetDefault("apidocs.host", "localhost:4869")
	viper.SetDefault("apidocs.basepath", "/")
	viper.SetDefault("apidocs.schemes", []string{"http"})

	viper.SetDefault("basic_auth.username", "admin")
	viper.SetDefault("basic_auth.password", "admin")
}
