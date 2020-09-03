package goutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRequestHTTPHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/api/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "get",
		})
	})

	getrecorder, err := RequestHTTPHandler(app, "GET", "/api/get", nil, nil)
	if err != nil {
		t.Error(err)
	}
	if getrecorder.Code != 200 {
		t.Error("code = ", getrecorder.Code)
	}

	jsonresp := struct {
		Data string `json:"data"`
	}{}

	json.Unmarshal(getrecorder.Body.Bytes(), &jsonresp)
	if jsonresp.Data != "get" {
		t.Error("data field value error:" + jsonresp.Data)
	}

	postdata := struct {
		Name string `json:"name"`
	}{}
	app.POST("/api/post", func(c *gin.Context) {
		c.ShouldBindJSON(&postdata)
		c.JSON(200, gin.H{
			"data": postdata.Name,
		})
	})
	body := []byte(`{"name": "axiaoxin"}`)
	header := map[string]string{
		"content-type": "application/json",
	}
	postrecorder, err := RequestHTTPHandler(app, "POST", "/api/post", body, header)
	if err != nil {
		t.Error(err)
	}
	if postrecorder.Code != 200 {
		t.Error("code = ", getrecorder.Code)
	}
	json.Unmarshal(postrecorder.Body.Bytes(), &jsonresp)
	if jsonresp.Data != "axiaoxin" {
		t.Error("data field value error:" + jsonresp.Data)
	}
}

func TestRequestHTTPHandlerFunc(t *testing.T) {
	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	}
	b, err := RequestHTTPHandlerFunc(f, "GET", nil, nil)
	if err != nil {
		t.Error(err)
	}
	if string(b) != "Hello!" {
		t.Error(string(b))
	}
}
