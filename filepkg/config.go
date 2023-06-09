package filepkg

import (
	"context"
	"github.com/BreezeHubs/go-pkg/timepkg"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"time"
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

func LoadYamlFileAndListen(file string, config any, changeFunc func(err error)) error {
	viper.SetConfigFile(file)

	if err := readConfig(&config); err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := readConfig(&config)
		if changeFunc != nil {
			changeFunc(err)
		}
	})
	return nil
}

func LoadYamlFileWithTickerAndListen(file string, config any, t time.Duration, readFunc func(err error)) {
	viper.SetConfigFile(file)

	tc := timepkg.TickerTaskWithChannel(context.Background(), func(ctx context.Context) error {
		return readConfig(&config)
	}, t)

	go func() {
		for {
			select {
			case err := <-tc.C:
				if readFunc != nil {
					readFunc(err)
				}
			}
		}
	}()
}

var readConfigLock sync.Mutex

func readConfig(config any) error {
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "Fatal error config file")
	}

	readConfigLock.Lock()
	defer readConfigLock.Unlock()

	if err := viper.Unmarshal(&config); err != nil {
		return errors.Wrap(err, "Unable to decode into struct")
	}
	return nil
}
