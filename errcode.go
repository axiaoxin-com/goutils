// 错误码封装定义

package goutils

import "strings"

// ErrCode 错误码结构体
type ErrCode struct {
	code interface{}
	msg  string
	errs []error
}

// NewErrCode 创建错误码,code可以是任意类型
func NewErrCode(code interface{}, msg string, errs ...error) *ErrCode {
	return &ErrCode{
		code: code,
		msg:  msg,
		errs: errs,
	}
}

// Decode 返回结构体里面的字段
func (c *ErrCode) Decode() (interface{}, string, []error) {
	return c.code, c.msg, c.errs
}

// Error 拼接所有err信息并返回
func (c *ErrCode) Error() string {
	errs := []string{}
	for _, err := range c.errs {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, "; ")
}

// Code 返回code
func (c *ErrCode) Code() interface{} {
	return c.code
}

// Msg 返回msg描述
func (c *ErrCode) Msg() string {
	return c.msg
}

// Errs 返回error列表
func (c *ErrCode) Errs() []error {
	return c.errs
}

// SetMsg 更新msg字段返回新的对象避免覆盖原始对象
func (c *ErrCode) SetMsg(msg string) *ErrCode {
	nc := *c
	nc.msg = msg
	return &nc
}

// AppendError 添加err到errs中
func (c *ErrCode) AppendError(errs ...error) *ErrCode {
	nc := *c
	nc.errs = append(nc.errs, errs...)
	return &nc
}
