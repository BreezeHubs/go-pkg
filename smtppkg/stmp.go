package smtppkg

import "gopkg.in/gomail.v2"

type QQEmail struct {
	from, name, code string
	conn             *gomail.Message
}

func NewQQEmail(from, name, code string) *QQEmail {
	m := gomail.NewMessage()

	m.SetHeader("From", from)            //发送人
	m.SetAddressHeader("Cc", from, name) //抄送人

	return &QQEmail{
		from: from,
		code: code,
		conn: m,
	}
}

func (e QQEmail) Set(subject, content string, to []string) error {
	//接收人
	e.conn.SetHeader("To", to...)
	//主题
	e.conn.SetHeader("Subject", subject)
	//内容
	e.conn.SetBody("text/html", content)
	//附件
	//m.Attach("./myIpPic.png")
	return e.send()
}

type FileSetting = gomail.FileSetting

func (e QQEmail) SetWithFile(subject, content string, to []string, filename string, settings ...FileSetting) error {
	//接收人
	e.conn.SetHeader("To", to...)
	//主题
	e.conn.SetHeader("Subject", subject)
	//内容
	e.conn.SetBody("text/html", content)
	//附件
	e.conn.Attach(filename, settings...)
	return e.send()
}

func (e *QQEmail) send() error {
	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, e.from, e.code)

	// 发送邮件
	return d.DialAndSend(e.conn)
}
