package game

import (
	"strconv"
	"sync"
)

type BulletFactory struct {
	bullets map[string]Bullet
	lock    sync.RWMutex
}

func NewBulletFactory() *BulletFactory {
	return &BulletFactory{bullets: make(map[string]Bullet)}
}

func (f *BulletFactory) GetBullet(color string, size int) Bullet {
	key := color + "_" + strconv.Itoa(size)
	f.lock.RLock()
	if bullet, ok := f.bullets[key]; ok {
		f.lock.RUnlock()
		return bullet
	}
	f.lock.RUnlock()

	f.lock.Lock()
	defer f.lock.Unlock()
	if bullet, ok := f.bullets[key]; ok {
		return bullet
	}
	bullet := NewDefaultBullet(color, size)
	f.bullets[key] = bullet
	return bullet
}
