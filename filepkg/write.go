package filepkg

import (
	"os"
	"path"
	"unsafe"
)

func WriteFileString(filePath string, content string) (int, error) {
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return 0, err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.Write(*(*[]byte)(unsafe.Pointer(&content)))
}
