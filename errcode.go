// 错误码封装定义

package goutils

import "strings"

// ErrCode 错误码结构体
type ErrCode struct {
	code interface{}
	msg  string
	errs []error
}

// NewErrCode 创建错误码,code 可以是任意类型
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

// Error 拼接所有 err 信息并返回
func (c *ErrCode) Error() string {
	errs := []string{}
	for _, err := range c.errs {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, "; ")
}

// Code 返回 code
func (c *ErrCode) Code() interface{} {
	return c.code
}

// Msg 返回 msg 描述
func (c *ErrCode) Msg() string {
	return c.msg
}

// Errs 返回 error 列表
func (c *ErrCode) Errs() []error {
	return c.errs
}

// SetMsg 更新 msg 字段
func (c *ErrCode) SetMsg(msg string) *ErrCode {
	// 返回新的对象避免覆盖原始对象
	nc := *c
	nc.msg = msg
	return &nc
}

// AppendMsg 追加 msg 字段
func (c *ErrCode) AppendMsg(msg string) *ErrCode {
	// 返回新的对象避免覆盖原始对象
	nc := *c
	nc.msg += ":" + msg
	return &nc
}

// AppendError 添加 err 到 errs 中
func (c *ErrCode) AppendError(errs ...error) *ErrCode {
	nc := *c
	nc.errs = append(nc.errs, errs...)
	return &nc
}
