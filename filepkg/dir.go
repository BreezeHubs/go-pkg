package filepkg

import (
	"errors"
	"os"
)

func DirList(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return nil, errors.New("dirPath does not exist")
	}

	dirEntry, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0, len(dirEntry))
	for _, entry := range dirEntry {
		if entry.IsDir() {
			name := entry.Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}
	return ret, nil
}
