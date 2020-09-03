package goutils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

// RequestHTTPHandler 根据参数请求传入的 http.Handler 对应的 path 接口，用于接口测试
// 返回 ResponseRecorder: https://golang.org/pkg/net/http/httptest/#ResponseRecorder
func RequestHTTPHandler(app http.Handler, method, path string, body []byte, header map[string]string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(strings.ToUpper(method), path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	rsp := httptest.NewRecorder()
	app.ServeHTTP(rsp, req)
	return rsp, nil
}

// RequestHTTPHandlerFunc 根据参数请求传入的 http.HandlerFunc 返回请求处理结果 body ，用于接口测试
func RequestHTTPHandlerFunc(f http.HandlerFunc, method string, body []byte, header map[string]string) ([]byte, error) {
	server := httptest.NewServer(http.HandlerFunc(f))
	defer server.Close()

	client := &http.Client{}
	req, err := http.NewRequest(strings.ToUpper(method), server.URL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return rspBody, nil
}
