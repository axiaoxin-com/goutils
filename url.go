package goutils

import (
	"fmt"
	"net/url"
	"path"
)

// AddQueryParam 向已有URL中追加查询参数
func AddQueryParam(baseURL, paramKey, paramValue string) (string, error) {
	// 解析原始URL
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	// 获取现有查询参数
	query := parsedURL.Query()

	// 添加新的查询参数
	query.Set(paramKey, paramValue)

	// 将新的查询参数设置回URL中
	parsedURL.RawQuery = query.Encode()

	// 生成并返回最终的URL
	return parsedURL.String(), nil
}

// JoinPath joins either filesystem paths or URL paths.
func JoinPath(base string, segments ...string) string {
	// Check if the base is a valid URL
	u, err := url.Parse(base)
	if err != nil || u.Scheme == "" || u.Host == "" {
		// Treat it as a filesystem path
		return path.Join(append([]string{base}, segments...)...)
	}

	// If it's a URL, ensure the segments are joined to the path part of the URL
	u.Path = path.Join(append([]string{u.Path}, segments...)...)
	return u.String()
}
