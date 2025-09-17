package sendmessage

import "testing"

func TestSendMessage(t *testing.T) {
	///var emailSender MessageSender
	emailSender := &EmailSender{}
	smsSender := &SmsSender{}
	wechatSender := &WechatSender{}

	senderService := NewSendService(emailSender)
	senderService.Register(smsSender)
	senderService.Register(wechatSender)

	err := senderService.Send("sms", "中奖了", "123", "13813813813")
	if err != nil {
		t.Error(err)
	}

}
