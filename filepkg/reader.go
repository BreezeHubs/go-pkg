package filepkg

import (
	"bufio"
	"os"
	"strings"
	"time"
)

type Reader struct {
	filePath string        // 配置的日志路径
	stream   chan string   // 同步日志的chan
	close    chan struct{} // 关闭的chan
	file     *os.File      // 文件读取对象
}

func NewReader(filePath string, stream chan string) (*Reader, <-chan struct{}, error) {
	r := &Reader{
		filePath: filePath,
		stream:   stream,
		close:    make(chan struct{}),
	}
	return r, r.close, r.openFile()
}

func (r *Reader) openFile() error {
	// 打开文件
	file, err := os.Open(r.filePath)
	if err != nil {
		return err
	}

	r.file = file
	return nil
}

func (r *Reader) Start(fn ...func(closeChan <-chan struct{}, readLine int64)) error {
	var readLine int64

	closeChan := make(chan struct{})
	for _, f := range fn {
		go f(closeChan, readLine)
	}

	// 获取文件信息
	fileInfo, err := r.file.Stat()
	if err != nil {
		return err
	}

	// 从文件结尾开始读取
	offset := fileInfo.Size()
	reader := bufio.NewReader(r.file)

	// 循环读取文件内容并发送到通道中
	for {
		// 读取行数据
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			// 如果是EOF则等待1秒钟后重新读取，否则退出循环
			return err
		}
		if line == "" {
			time.Sleep(100 * time.Microsecond)
			continue
		}

		// 发送到通道中
		r.stream <- strings.TrimRight(line, "\n")
		offset += int64(len(line))

		// 检查文件是否有新内容
		fileInfo, err = r.file.Stat()
		if err != nil {
			return err
		}
		if fileInfo.Size() < offset {
			offset = fileInfo.Size()
			if _, err = r.file.Seek(offset, 0); err != nil {
				return err
			}
			reader = bufio.NewReader(r.file)
		}
	}
	close(closeChan)
	return nil
}

func (r *Reader) Stop() {
	r.file.Close()
	close(r.close)
}
