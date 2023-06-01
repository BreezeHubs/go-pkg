package typexpkg

import (
	"unsafe"
)

func BytesToString(b []byte) string {
	if b == nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}
