package ginpkg

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// FilterBindErr 自定义错误消息
func FilterBindErr(errs error, r any) error {
	if errs == nil {
		return nil
	}

	vErrs := errs.(validator.ValidationErrors)
	s := reflect.TypeOf(r)
	for _, fieldError := range vErrs {
		filed, _ := s.FieldByName(fieldError.Field()) //用反射获取参数名

		//获取对应binding得错误消息 tag：错误类型_err，如：required_err:"userId不能为空"`
		if errTagText := filed.Tag.Get(fieldError.Tag() + "_err"); errTagText != "" {
			return errors.New(errTagText)
		}

		//通用错误类型 tag：err，如：err:"userId错误"`
		if errText := filed.Tag.Get("err"); errText != "" {
			return errors.New(errText)
		}

		//无法匹配到错误，如：user的格式需遵守: required
		return errors.New(fieldError.Field() + "的格式需遵守: " + fieldError.Tag())
	}
	return nil
}
