package goutils

import (
	"fmt"
	"net/url"
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
	query.Add(paramKey, paramValue)

	// 将新的查询参数设置回URL中
	parsedURL.RawQuery = query.Encode()

	// 生成并返回最终的URL
	return parsedURL.String(), nil
}
