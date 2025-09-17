package sendmessage

import "fmt"

type MessageSender interface {
	Send(subject, body, target string) error
	SendType() string
}

type EmailSender struct {
}

func (e *EmailSender) Send(subject, body, target string) error {
	fmt.Printf("send a mail subject %s body %s to %s\n", subject, body, target)
	return nil
}

func (e *EmailSender) SendType() string {
	return "email"
}

type SmsSender struct {
}

func (s *SmsSender) Send(subject, body, target string) error {
	fmt.Printf("send a sms subject %s body %s to %s\n", subject, body, target)
	return nil
}

func (s *SmsSender) SendType() string {
	return "sms"
}

type WechatSender struct {
}

func (w *WechatSender) Send(subject, body, target string) error {
	fmt.Printf("send a wechat subject %s body %s to %s\n", subject, body, target)
	return nil
}
func (w *WechatSender) SendType() string {
	return "wechat"
}
