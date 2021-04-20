package goutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// NewHTTPMultipartReq 根据参数创建form-data请求
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
	return req, nil
}

// HTTPPOST 发送 http post 请求
func HTTPPOST(ctx context.Context, cli *http.Client, apiurl string, req *http.Request, rspPointer interface{}) error {
	resp, err := cli.Do(req)
	if err != nil {
		return errors.Wrap(err, "POST do request error")
	}
	defer resp.Body.Close()

	rspbuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, fmt.Sprint("POST read resp body error, resp.Body:", resp.Body))
	}
	if err := json.Unmarshal(rspbuf, rspPointer); err != nil {
		return errors.Wrap(err, "POST json unmarshal result error")
	}
	return nil
}

// NewHTTPGetURLWithQueryString 创建带querystring的http get 请求url
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

// HTTPGET 发送 http get 请求
func HTTPGET(ctx context.Context, cli *http.Client, apiurl string, rspPointer interface{}) error {
	resp, err := cli.Get(apiurl)
	if err != nil {
		return errors.Wrap(err, "Get request error")
	}
	defer resp.Body.Close()

	rspbuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, fmt.Sprint("GET read resp body error, resp.Body:", resp.Body))
	}
	if err := json.Unmarshal(rspbuf, rspPointer); err != nil {
		return errors.Wrap(err, "GET json unmarshal result error")
	}
	return nil
}
