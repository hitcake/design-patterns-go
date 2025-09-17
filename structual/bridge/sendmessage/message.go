package sendmessage

import "fmt"

type Message struct {
	sender  MessageSender
	subject string
	content string
	target  string
}

func (m *Message) SetSender(s MessageSender) {
	m.sender = s
}
func (m *Message) Send() error {
	return m.sender.Send(m.subject, m.content, m.target)
}

type SystemAlertMessage struct {
	Message
}

func (m *SystemAlertMessage) Send() error {
	m.subject = fmt.Sprintf("[系统报警] %s", m.subject)
	return m.Send()
}

type MarketingMessage struct {
	Message
}

func (m *MarketingMessage) Send() error {
	m.subject = fmt.Sprintf("%s 促销活动", m.subject)
	m.content = fmt.Sprintf("尊敬的客户 %s 了解更多", m.content)
	return m.Send()
}
