package light

type Command interface {
	Execute()
	Undo()
}

type TurnOnCommand struct {
	light *Light
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}
func (c *TurnOnCommand) Undo() {
	c.light.TurnOff()
}

type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}
func (c *TurnOffCommand) Undo() {
	c.light.TurnOn()
}
