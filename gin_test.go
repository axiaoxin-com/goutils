package goutils

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewGinEngine(t *testing.T) {
	GinPprofURLPath = "/x/test/pprof"
	router := NewGinEngine(gin.DebugMode, true)
	if router == nil {
		t.Error("new a nil gin engine")
	}
}
