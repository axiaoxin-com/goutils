[#](#) goutils

[![Build Status](https://travis-ci.org/axiaoxin-com/goutils.svg?branch=master)](https://travis-ci.org/axiaoxin-com/goutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/axiaoxin-com/goutils)](https://goreportcard.com/report/github.com/axiaoxin-com/goutils)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/927424c522de4548afa6b53cffd2e154)](https://www.codacy.com/gh/axiaoxin-com/goutils?utm_source=github.com&utm_medium=referral&utm_content=axiaoxin-com/goutils&utm_campaign=Badge_Grade)

Golang 通用类函数工具包

在线文档：<https://godoc.org/github.com/axiaoxin-com/goutils>

## 功能概览


### Redis 相关

- [redis 常用操作封装（包括但不限于：创建各类连接客户端实例，根据 viper 配置直接获取对应实例）](./redis.go)

### 文件相关

- [复制文件: CopyFile](./file.go)

### 单元测试相关

- [模拟 http 请求对应的 http.Handler](./httptest.go)

### IP 相关

- [获取当前 IP: GetLocalIP](./ip.go)

### 时间相关

- [自定义结构体序列化为 JSON 时时间类型字段的格式](./jsontime.go)

### slice 相关

- [根据下标删除 string slice 中的元素: RemoveStringSliceItemByIndex](./slice.go)
- [判断两个 string slice 是否相同: IsEqualStringSlice](./slice.go)
- [float64 slice 按指定大小进行切块: ChunkFloat64Slice](./slice.go)
- [判断字符串是否在给定的字符串列表中: IsStrInSlice](./slice.go)
- [判断 int 是否在给定的 int 列表中: IsIntInSlice](./slice.go)

### string 相关

- [删除字符串中所有的空白符: RemoveAllWhitespace](./string.go)
- [反转字符串: ReverseString](./string.go)
- [删除字符串中的重复空白符: RemoveDuplicateWhitespace](./string.go)

### 配置相关

- [根据配置文件路径和名称初始化 viper](./viper.go)

### URL 相关

- [生成 url key: URLKey](./url.go)

### Struct 相关

- [将结构体指针对象转换为 url.Values: StructToURLValues](./struct.go)
- [将结构体指针对象转换为 map[string]interface{}: StructToMap](./struct.go)
- [获取结构体指针对象 tag 列表: StructTagList](./struct.go)

### HTTP 请求相关

- [根据参数创建 json post 请求: NewHTTPJSONReq](./http.go)
- [根据参数创建 form-data post 请求: NewHTTPMultipartReq](./http.go)
- [发送 http post 请求: HTTPPOSTRaw](./http.go)
- [发送 http post 请求并将结果进行 json unmarsal: HTTPPOST](./http.go)
- [创建带 querystring 的 http get 请求 url: NewHTTPGetURLWithQueryString](./http.go)
- [发送 http get 请求: HTTPGETRaw](./http.go)
- [发送 http get 请求并将结果进行 json unmarsal: HTTPGET](./http.go)

### 时间相关

- [Unix 时间戳规整： UnixTImestampTrim](./time.go)
- [字符串时间转时间对象： StrToTime](./time.go)
- [获取最近一个交易日日期： GetLatestTradingDay](./time.go)
- [当期时间是否是交易日： IsTradingDay](./time.go)

### 统计

- [求 float64 列表均值: AvgFloat64](./statistics.go)
- [求 float64 列表方差: VarianceFloat64](./statistics.go)
- [求 float64 列表标准差: StdDeviationFloat64](./statistics.go)
- [求 float64 列表中位数: MidValueFloat64](./statistics.go)

### 其他分类

- [错误码结构体对象封装](./errcode.go)
- [hashids: 生成可相互转换的数字类型的 ID 与随机字符串 ID](./hashids.go)
- [分页计算](./pagination.go)
- [validator 参数验证错误信息自定义](./validator.go)
- [封装 go-cache 的 Get 方法支持直接获取具体类型: GetGoCache](./gocache.go)
- [验证邮箱地址是否合法](./email.go)
- [密码加密、校验](./password.go)
