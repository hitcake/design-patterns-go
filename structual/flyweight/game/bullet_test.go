package game

import "testing"

func TestBullet(t *testing.T) {
	factory := NewBulletFactory()
	bulletSlice := make([]Bullet, 10)
	for i := 0; i < len(bulletSlice); i++ {
		if i%2 == 0 {
			bulletSlice[i] = factory.GetBullet("red", 10)
		} else {
			bulletSlice[i] = factory.GetBullet("green", 10)
		}
	}
	for index, bullet := range bulletSlice {
		for i := 0; i < 10; i++ {
			bullet.Show(index, index+i)
		}
	}

}
