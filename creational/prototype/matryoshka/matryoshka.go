package matryoshka

import (
	"errors"
	"fmt"
)

type Matryoshka interface {
	Clone() Matryoshka
	GetSize() int
	GetName() string
	AddChild(m Matryoshka) error
	GetChild() Matryoshka
	Display()
}

type DefaultMatryoshka struct {
	Size  int
	Name  string
	child Matryoshka
}

func NewDefaultMatryoshka(name string, size int) Matryoshka {
	return &DefaultMatryoshka{Size: size, Name: name}
}

func (m *DefaultMatryoshka) Clone() Matryoshka {
	newSize := m.Size - 1
	if newSize <= 0 {
		return nil
	}
	return NewDefaultMatryoshka(m.Name, newSize)
}
func (m *DefaultMatryoshka) GetSize() int {
	return m.Size
}
func (m *DefaultMatryoshka) GetName() string {
	return m.Name
}
func (m *DefaultMatryoshka) AddChild(child Matryoshka) error {
	if m.child != nil {
		return errors.New("child already exists")
	}
	if child == nil {
		return errors.New("child is nil")
	}
	if m.Size > child.GetSize() {
		m.child = child
		return nil
	}
	return errors.New("child size is too big")
}
func (m *DefaultMatryoshka) GetChild() Matryoshka {
	return m.child
}
func (m *DefaultMatryoshka) Display() {
	var p Matryoshka = m
	for p != nil {
		fmt.Printf("%s 大小 %d\n", p.GetName(), p.GetSize())
		p = p.GetChild()
	}
}
