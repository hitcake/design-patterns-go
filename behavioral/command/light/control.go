package light

// Invoker
type Control struct {
	buttonOn  Command
	buttonOff Command
	//保存上一次按了哪个键
	historyButton Command
}

func NewControl(l *Light) *Control {
	return &Control{buttonOn: &TurnOnCommand{l}, buttonOff: &TurnOffCommand{l}}
}
func (c *Control) PressOn() {
	c.buttonOn.Execute()
	c.historyButton = c.buttonOn
}
func (c *Control) PressOff() {
	c.buttonOff.Execute()
	c.historyButton = c.buttonOff
}
func (c *Control) PressBack() {
	if c.historyButton != nil {
		c.historyButton.Undo()
	}
	c.historyButton = nil
}
