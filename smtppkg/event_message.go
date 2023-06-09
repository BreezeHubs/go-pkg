package smtppkg

import (
	"fmt"
	"github.com/BreezeHubs/go-pkg/filepkg"
)

type EventMessage struct {
	From     string `json:"from"`      // 来源
	Type     string `json:"type"`      // 类型
	Tag      string `json:"tag"`       // 标签
	Event    string `json:"event"`     // 事件
	Content  string `json:"content"`   // 内容
	Time     string `json:"time"`      // 时间
	SendTime string `json:"send_time"` // 发送时间
}

func (msg *EventMessage) GetDefaultContent() string {
	return fmt.Sprintf(
		"<p>来源：<span style='font-weight: bold;'>%s</span></p>"+
			"<p>类型：<span style=''>%s</span></p>"+
			"<p>标签：<span>%s</span></p>"+
			"<p>事件：<span>%s</span></p>"+
			"<p>内容：<span>%s</span></p>"+
			"<p>事件时间：<span>%s</span></p>"+
			"<p>发送时间：<span>%s</span></p>",
		msg.From, msg.Type, msg.Tag, msg.Event, msg.Content, msg.Time, msg.SendTime,
	)
}

func (msg *EventMessage) GetContentByFile(file, typeColor string) string {
	s, err := filepkg.ReadString(file)
	if err != nil {
		return ""
	}
	return fmt.Sprintf(s, msg.From, typeColor, msg.Type, msg.Tag, msg.Event, msg.Content, msg.Time, msg.SendTime)
}
