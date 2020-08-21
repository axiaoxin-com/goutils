// web 框架 gin 相关封装

package goutils

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	// GinPprofURLPath 设置 gin 中的 pprof url 注册路径，可以通过外部修改
	GinPprofURLPath = "/x/pprof"
)

// NewGinEngine 根据参数创建 gin 的 router engine
// mode gin.ReleaseMode gin.TestMode gin.DebugMode
// registerPprof 是否注册 pprof
// middlewares 需要使用到的中间件列表，默认不为 engine 添加任何中间件
func NewGinEngine(mode string, registerPprof bool, middlewares ...gin.HandlerFunc) *gin.Engine {
	// set gin mode
	if mode == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}
	gin.SetMode(mode)

	engine := gin.New()

	// use middlewares
	for _, middleware := range middlewares {
		engine.Use(middleware)
	}

	if registerPprof {
		pprof.Register(engine, GinPprofURLPath)
	}
	return engine
}
