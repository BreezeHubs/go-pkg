package ginpkg

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func Recover(isPrintErr, isStack bool, f func(c *gin.Context, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if isStack {
					//打印错误堆栈信息
					log.Printf("panic: %#v\n", r)
					debug.PrintStack()
				}

				var errStr = "服务器内部错误"
				if isPrintErr {
					errStr += fmt.Sprintf(" panic: %#v\n", r)
				}

				//封装通用json返回
				f(c, errors.New(errStr))
			}
		}()
		c.Next()
	}
}
