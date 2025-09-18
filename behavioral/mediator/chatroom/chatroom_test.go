package chatroom

import "testing"

func TestChat(t *testing.T) {
	boss := NewUser("boss")
	manager := NewUser("manager")
	designer := NewUser("设计师老刘")
	tester := NewUser("测试老郑")
	engineer := NewUser("研发老张")
	project001Room := NewChatRoom("项目X研发群")
	project001Room.AddUser(boss, manager, designer, tester, engineer)
	boss.SendMsg(project001Room, "明早10点，在山水阁会议室开个会，碰一下目前的进度情况")
	manager.SendMsg(project001Room, "收到")
	designer.SendMsg(project001Room, "收到")
	tester.SendMsg(project001Room, "收到")
	engineer.SendMsg(project001Room, "收到")
}
