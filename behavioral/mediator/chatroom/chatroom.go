package chatroom

type ChatRoom struct {
	roomName string
	users    []*User
}

func NewChatRoom(roomName string) *ChatRoom {
	return &ChatRoom{roomName: roomName, users: make([]*User, 0)}
}

func (c *ChatRoom) AddUser(user ...*User) {
	c.users = append(c.users, user...)
}
func (c *ChatRoom) RemoveUser(user *User) {
	for i, u := range c.users {
		if u == user {
			c.users = append(c.users[:i], c.users[i+1:]...)
			return
		}
	}
}
func (c *ChatRoom) SendMessage(speaker *User, message string) {
	for _, user := range c.users {
		if user != speaker {
			user.ReceiveMsg(c, speaker, message)
		}
	}
}
