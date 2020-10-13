// gin binding validator

// Usage
// func init() {
//     binding.Validator = &GinStructValidator{}
// }

package goutils

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// GinStructValidator 自定义参数 binding 验证错误信息输出格式
type GinStructValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &GinStructValidator{}

// ValidatorTagName 结构体 validator 的 tag 名
var ValidatorTagName = "binding"

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *GinStructValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if errs := v.validate.Struct(obj); errs != nil {
			var errmsg string
			for _, e := range errs.(validator.ValidationErrors) {
				errmsg += fmt.Sprintf("%s;", ValidationErrorToText(e))
			}
			return errors.New(errmsg)
		}
	}
	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validator.v8
func (v *GinStructValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *GinStructValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName(ValidatorTagName)
	})
}

// ValidationErrorToText error msg for human
func ValidationErrorToText(e validator.FieldError) string {
	switch e.ActualTag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be more than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be more than %s", e.Field(), e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", e.Field(), e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of:[%s]", e.Field(), e.Param())
	case "datetime":
		return fmt.Sprintf("%s must use format as %s", e.Field(), e.Param())
	}
	return fmt.Sprintf("%s is not valid", e.Field())
}
