package light

import "fmt"

// 接受者
type Light struct {
	name string
	on   bool
}

func NewLight(name string) *Light {
	return &Light{name: name, on: false}
}

func (l *Light) TurnOn() {
	l.on = true
	fmt.Println("light is on")
}

func (l *Light) TurnOff() {
	l.on = false
	fmt.Println("light is off")
}
