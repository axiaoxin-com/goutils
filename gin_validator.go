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
				errmsg += fmt.Sprintf("%s;", GinValidationErrorToText(e))
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
		v.validate.SetTagName("binding")
	})
}

// GinValidationErrorToText error msg for human
func GinValidationErrorToText(e validator.FieldError) string {
	switch e.ActualTag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be bigger than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be bigger than %s", e.Field(), e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", e.Field(), e.Param())
	}
	return fmt.Sprintf("%s is not valid", e.Field())
}

// Usage
// func init() {
//     binding.Validator = &GinStructValidator{}
// }
