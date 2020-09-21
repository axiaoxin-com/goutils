# goutils

[![Build Status](https://travis-ci.org/axiaoxin-com/goutils.svg?branch=master)](https://travis-ci.org/axiaoxin-com/goutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/axiaoxin-com/goutils)](https://goreportcard.com/report/github.com/axiaoxin-com/goutils)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/927424c522de4548afa6b53cffd2e154)](https://www.codacy.com/gh/axiaoxin-com/goutils?utm_source=github.com&utm_medium=referral&utm_content=axiaoxin-com/goutils&utm_campaign=Badge_Grade)

Golang 通用类函数工具包

在线文档：<https://godoc.org/github.com/axiaoxin-com/goutils>

## 功能概览

### 数据库相关

- [根据配置信息生成不同数据库的 DSN 字符串](./dbconfig.go)
- [gorm 常用操作封装（包括但不限于：创建各类 db 实例，根据 viper 配置直接获取各类 db 实例）](./gormdb.go)
- [sqlx 常用操作封装（包括但不限于：创建各类 db 实例，根据 viper 配置直接获取各类 db 实例）](./sqlxdb.go)

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

### string 相关

- [删除字符串中所有的空白符: RemoveAllWhitespace](./string.go)
- [反转字符串: ReverseString](./string.go)

### 配置相关

- [根据配置文件路径和名称初始化 viper](./viper.go)

### URL 相关

- [将结构体指针对象转换为 url.Values: StructToURLValues](./struct.go)
- [生成 url key: URLKey](./url.go)

### 其他分类

- [错误码结构体对象封装](./errcode.go)
- [hashids: 生成可相互转换的数字类型的 ID 与随机字符串 ID](./hashids.go)
- [分页计算](./pagination.go)
- [gin binding 参数验证错误信息自定义](./gin_validator.go)
