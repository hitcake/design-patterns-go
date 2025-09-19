package light

import "testing"

func TestCommand(t *testing.T) {
	light := NewLight("主卧")
	control := NewControl(light)

	control.PressOn()
	control.PressOff()
	control.PressOn()
	control.PressBack()
}
