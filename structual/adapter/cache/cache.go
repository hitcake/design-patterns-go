package cache

// Cache /* v是string类型，不是很实用，仅作为适配器的参考
type Cache interface {
	Put(k string, v string) error
	Get(k string) (string, error)
	Remove(k string)
}

type DefaultCache struct {
	cache map[string]string
}

func NewDefaultCache() *DefaultCache {
	return &DefaultCache{make(map[string]string)}
}

func (c *DefaultCache) Put(k string, v string) error {
	c.cache[k] = v
	return nil
}

func (c *DefaultCache) Get(k string) (string, error) {
	return c.cache[k], nil
}
func (c *DefaultCache) Remove(k string) {
	delete(c.cache, k)
}
