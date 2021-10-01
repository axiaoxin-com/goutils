package goutils

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// NewHTTPJSONReq 根据参数创建 json 请求
func NewHTTPJSONReq(ctx context.Context, apiurl string, reqData interface{}) (*http.Request, error) {
	reqbuf, err := jsoniter.Marshal(reqData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprint("json marshal error"))
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiurl, bytes.NewReader(reqbuf))
	if err != nil {
		return nil, errors.Wrap(err, "Post NewRequestWithContext")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "app-iphone-client-iPhone11,"+fmt.Sprint(time.Now().Unix()))
	return req, nil
}

// NewHTTPMultipartReq 根据参数创建 form-data 请求
func NewHTTPMultipartReq(ctx context.Context, apiurl string, reqData map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range reqData {
		if err := writer.WriteField(k, v); err != nil {
			return nil, errors.Wrap(err, "WriteField error")
		}
	}
	if err := writer.Close(); err != nil {
		return nil, errors.Wrap(err, "Writer close error")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiurl, body)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequestWithContext error")
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "app-iphone-client-iPhone11,"+fmt.Sprint(time.Now().Unix()))
	return req, nil
}

// HTTPPOST 发送 http post 请求并将结果 json unmarshal 到 rspPointer
func HTTPPOST(ctx context.Context, cli *http.Client, req *http.Request, rspPointer interface{}) error {
	rspbuf, err := HTTPPOSTRaw(ctx, cli, req)
	if err != nil {
		return err
	}
	if err := jsoniter.Unmarshal(rspbuf, rspPointer); err != nil {
		return errors.Wrap(err, fmt.Sprintf("json unmarshal result error, rspbuf:%s", string(rspbuf)))
	}
	return nil
}

// HTTPPOSTRaw 发送 http post 请求
func HTTPPOSTRaw(ctx context.Context, cli *http.Client, req *http.Request) ([]byte, error) {
	resp, err := cli.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "POST do request error")
	}
	defer resp.Body.Close()

	rspbuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprint("read resp body error, resp.Body:", resp.Body))
	}
	return rspbuf, nil
}

// NewHTTPGetURLWithQueryString 创建带 querystring 的 http get 请求 url
func NewHTTPGetURLWithQueryString(ctx context.Context, apiurl string, params map[string]string) (string, error) {
	value := url.Values{}
	u, err := url.Parse(apiurl)
	if err != nil {
		return "", err
	}
	for k, v := range params {
		value.Set(k, v)
	}
	u.RawQuery = value.Encode()
	apiurl = u.String()
	return apiurl, nil
}

// HTTPGET 发送 http get 请求并将返回结果 json unmarshal 到 rspPointer
func HTTPGET(ctx context.Context, cli *http.Client, apiurl string, header map[string]string, rspPointer interface{}) error {
	rspbuf, err := HTTPGETRaw(ctx, cli, apiurl, header)
	if err != nil {
		return err
	}
	if err := jsoniter.Unmarshal(rspbuf, rspPointer); err != nil {
		return errors.Wrap(err, fmt.Sprintf("json unmarshal result error, rspbuf:%s", string(rspbuf)))
	}
	return nil
}

// HTTPGETRaw 发送 http get 请求
func HTTPGETRaw(ctx context.Context, cli *http.Client, apiurl string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiurl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequestWithContext error")
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Get request error")
	}
	defer resp.Body.Close()

	rspbuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprint("read resp body error, resp.Body:", resp.Body))
	}
	return rspbuf, nil
}
