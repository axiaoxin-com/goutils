package goutils

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/axiaoxin-com/logging"
	"github.com/spf13/viper"
)

// RunWebApp 以 viper 加载的 webapp 配置启动运行 http.Handler 的 app
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
