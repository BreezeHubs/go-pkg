package typexpkg

import "unsafe"

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
