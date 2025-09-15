package matryoshka

import (
	"fmt"
	"testing"
)

func TestMatryoshka(t *testing.T) {
	prototype := NewDefaultMatryoshka("红色套娃", 6)
	p := prototype
	for {
		child := p.Clone()
		if child == nil {
			fmt.Println("done")
			break
		}
		_ = p.AddChild(child)
		p = child
	}
	prototype.Display()
}
