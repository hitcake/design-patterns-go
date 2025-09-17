package game

import "fmt"

type Bullet interface {
	GetColor() string
	GetSize() int
	Show(x, y int)
}

type DefaultBullet struct {
	color string
	size  int
}

func NewDefaultBullet(color string, size int) *DefaultBullet {
	fmt.Printf("create new bullet %s %d\n", color, size)
	return &DefaultBullet{color: color, size: size}
}
func (b *DefaultBullet) GetColor() string {
	return b.color
}
func (b *DefaultBullet) GetSize() int {
	return b.size
}
func (b *DefaultBullet) Show(x, y int) {
	fmt.Printf("bullet color %s size %d shoot (%d,%d)\n", b.GetColor(), b.size, x, y)
}
