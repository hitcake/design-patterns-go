package chatroom

import "fmt"

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name}
}
func (u *User) ReceiveMsg(room *ChatRoom, speaker *User, message string) {
	fmt.Printf("(%s) 来自群[%s]的消息 ", u.name, room.roomName)
	fmt.Printf("(%s) %s 说: %s\n", u.name, speaker.name, message)
}
func (u *User) SendMsg(room *ChatRoom, message string) {
	room.SendMessage(u, message)
}
