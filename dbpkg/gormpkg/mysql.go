package gormpkg

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"path"
	"time"
)

type Config struct {
	DSN              string `yaml:"dsn"`
	LogType          int8   `yaml:"log_type"`
	LogFile          string `yaml:"log_file"`
	SlowSqlThreshold int    `yaml:"slow_sql_threshold"`
	LogLevel         string `yaml:"log_level"`
}

type DB struct {
	Engine *gorm.DB
}

func NewMySQL(conf *Config) (*DB, error) {
	lw, err := loggerWriter(conf.LogType, conf.LogFile)
	if err != nil {
		return nil, err
	}

	gdb, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: conf.DSN, // DSN data source name
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
			NowFunc: func() time.Time {
				return time.Now().Local() // 更改创建时间使用的函数
			},
			PrepareStmt: true, // 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
			Logger: logger.New(lw,
				logger.Config{
					SlowThreshold:             time.Duration(conf.SlowSqlThreshold) * time.Second, // Slow SQL threshold
					LogLevel:                  loggerLevel(conf.LogLevel),                         // Log level
					IgnoreRecordNotFoundError: true,                                               // Ignore ErrRecordNotFound error for logger
				},
			),
		})
	if err != nil {
		return nil, err
	}
	return &DB{
		Engine: gdb,
	}, nil
}

func loggerLevel(logLevel string) logger.LogLevel {
	switch logLevel {
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	default:
		return logger.Silent
	}
}

func loggerWriter(logType int8, logFile string) (logger.Writer, error) {
	var lw logger.Writer
	if logType == 1 {
		if err := os.MkdirAll(path.Dir(logFile), os.ModePerm); err != nil {
			return nil, err
		}
		logFile, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open db log file")
		}
		lw = log.New(logFile, "\r\n", log.LstdFlags)
	} else {
		lw = log.New(os.Stdout, "\r\n", log.LstdFlags)
	}
	return lw, nil
}
