package goutils

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/axiaoxin-com/logging"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

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

// RunWebApp 以 viper 加载的 webapp 配置启动运行 http.Handler 的 app
// 注意：这里依赖 viper ，必须在外部先对 viper 配置进行加载
func RunWebApp(app http.Handler, routesRegister func(http.Handler)) {
	// 结束时关闭 db 连接
	defer CloseGormInstances()

	// 判断是否加载 viper 配置
	if !IsInitedViper() {
		panic("RunWebApp must init viper by config file first!")
	}

	// 注册 api 路由
	routesRegister(app)

	// 创建 server
	addr := viper.GetString("server.addr")
	readTimeout := viper.GetInt("server.read_timeout")
	writeTimeout := viper.GetInt("server.write_timeout")
	srv := &http.Server{
		Addr:         addr,
		Handler:      app,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	// 启动 http server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatal(nil, "Server start error:"+err.Error())
		}
	}()
	logging.Info(nil, "Server is listening and serving on "+srv.Addr)

	// 监听中断信号， WriteTimeout 时间后优雅关闭服务
	// syscall.SIGTERM 不带参数的 kill 命令
	// syscall.SIGINT ctrl-c kill -2
	// syscall.SIGKILL 是 kill -9 无法捕获这个信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Infof(nil, "Server will shutdown after %d seconds", writeTimeout)

	// 创建一个 context 用于通知 server 有 writeTimeout 秒的时间结束当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(writeTimeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Fatal(nil, "Server shutdown with error: "+err.Error())
	}
	logging.Info(nil, "Server shutdown")
}
