package netpkg

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"time"
)

type HeaderOption struct {
	Key   string
	Value string
}

var JsonHeaderOption = &HeaderOption{
	Key:   "Content-Type",
	Value: "application/json;charset=UTF-8",
}

type curl struct {
	req *http.Request
	err error
}

func Post(url string, data map[string]string) *curl {
	var (
		dataJson []byte    //json数据
		err      error     //error
		c        = &curl{} //请求结构体数据
	)

	//处理map成json数据
	if data != nil && len(data) > 0 {
		dataJson, err = jsoniter.Marshal(data)
		if err != nil {
			c.err = err
			return c
		}
	}

	//创建request请求
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(dataJson))
	if err != nil {
		c.err = err
		return c
	}

	//成功创建请求结构体数据
	c.req = request
	return c
}

func Get(url string) *curl {
	var c = &curl{} //请求结构体数据

	//创建request请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.err = err
		return c
	}

	//成功创建请求结构体数据
	c.req = request
	return c
}

func (c *curl) Do(headers map[string]string, t time.Duration) ([]byte, error) {
	//抛出错误
	if c.err != nil {
		return nil, c.err
	}

	//设置请求头
	if headers != nil && len(headers) > 0 {
		for headerKey, headerValue := range headers {
			c.req.Header.Set(headerKey, headerValue)
		}
	}

	client := http.Client{Timeout: t}
	resp, err := client.Do(c.req) //发送请求
	if err != nil {
		return nil, err
	}

	//返回数据
	return io.ReadAll(resp.Body)
}
