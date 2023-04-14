package filepkg

import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

func ReadStringTrim(filePath string) (string, error) {
	s, err := ReadString(filePath)
	return strings.TrimSpace(s), err
}

func ReadString(filePath string) (string, error) {
	bytes, err := ReadBytes(filePath)
	return *(*string)(unsafe.Pointer(&bytes)), err
}

func ReadBytes(filePath string) ([]byte, error) {
	if !IsExist(filePath) {
		return nil, fmt.Errorf("%s not exists", filePath)
	}

	if !IsFile(filePath) {
		return nil, fmt.Errorf("%s not file", filePath)
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
