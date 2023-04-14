package filepkg

import (
	"gopkg.in/yaml.v3"
	"os"
)

// LoadYamlFile 根据conf路径读取内容
func LoadYamlFile(filename string, config any) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return load(content, config)
}

// 根据io read读取配置文件后的字符串解析yaml
func load(s []byte, config any) error {
	return yaml.Unmarshal(s, config)
}
