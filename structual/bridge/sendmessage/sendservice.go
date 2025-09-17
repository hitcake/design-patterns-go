package sendmessage

type SendService struct {
	defaultSender MessageSender
	senders       map[string]MessageSender
}

func NewSendService(defaultSender MessageSender) *SendService {
	return &SendService{defaultSender: defaultSender, senders: make(map[string]MessageSender)}
}
func (s *SendService) Register(sender MessageSender) {
	s.senders[sender.SendType()] = sender
}

func (s *SendService) getSender(sendType string) MessageSender {
	if send, ok := s.senders[sendType]; ok {
		return send
	}
	return s.defaultSender
}

func (s *SendService) Send(sendType, subject, body, target string) error {
	sender := s.getSender(sendType)
	if sender == nil {
		sender = s.defaultSender
	}
	return sender.Send(subject, body, target)
}
