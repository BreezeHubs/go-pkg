package datapkg

import "reflect"

func IsEqual(a, b any) bool {
	// 如果a和b类型不同，则直接返回false
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	// 获取a和b的值
	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)

	// 如果a和b都是零值或nil，则返回true
	if va.IsZero() && vb.IsZero() || !va.IsValid() && !vb.IsValid() {
		return true
	}

	// 如果a和b是切片类型，则使用 reflect.DeepEqual() 函数比较它们是否相等
	if va.Type().Kind() == reflect.Slice {
		return reflect.DeepEqual(va.Interface(), vb.Interface())
	}

	// 如果a和b的类型是基本类型，则直接比较它们的值
	if va.Type().Kind() != reflect.Struct {
		return va.Interface() == vb.Interface()
	}

	// 遍历a和b的所有字段，递归比较嵌套结构体的字段
	for i := 0; i < va.NumField(); i++ {
		if !IsEqual(va.Field(i).Interface(), vb.Field(i).Interface()) {
			return false
		}
	}
	return true
}
