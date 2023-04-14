package filepkg

import "os"

func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

func IsFile(fp string) bool {
	f, e := os.Stat(fp)
	return e == nil || !f.IsDir()
}
