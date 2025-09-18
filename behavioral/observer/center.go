package observer

import "sync"

type Center struct {
	listenerMap  map[string][]Listener
	listenerLock sync.RWMutex
	serviceMap   map[string][]string
	servicesLock sync.RWMutex
}

func NewCenter() *Center {
	return &Center{listenerMap: make(map[string][]Listener), serviceMap: make(map[string][]string)}
}

func (c *Center) Register(service string, l Listener) {
	c.listenerLock.Lock()
	defer c.listenerLock.Unlock()

	if listeners, ok := c.listenerMap[service]; ok {
		listeners = append(listeners, l)
		c.listenerMap[service] = listeners
	} else {
		c.listenerMap[service] = []Listener{l}
	}
}
func (c *Center) UnRegister(service string, l Listener) {
	c.listenerLock.Lock()
	defer c.listenerLock.Unlock()
	if listeners, ok := c.listenerMap[service]; ok {
		for i := len(listeners) - 1; i >= 0; i-- {
			if l == listeners[i] {
				listeners = append(listeners[:i], listeners[i+1:]...)
				c.listenerMap[service] = listeners
				break
			}
		}
	}
}

func (c *Center) Up(service string, url string) {
	c.servicesLock.Lock()
	defer c.servicesLock.Unlock()
	if listeners, ok := c.serviceMap[service]; ok {
		listeners = append(listeners, url)
		c.serviceMap[service] = listeners
	} else {
		c.serviceMap[service] = []string{url}
	}
	go func() {
		c.listenerLock.RLock()
		defer c.listenerLock.RUnlock()
		for _, l := range c.listenerMap[service] {
			l.notify(url, "up")
		}
	}()
}

func (c *Center) Down(service string, url string) {
	c.servicesLock.Lock()
	defer c.servicesLock.Unlock()
	if urls, ok := c.serviceMap[service]; ok {
		for i := len(urls) - 1; i >= 0; i-- {
			if urls[i] == url {
				urls = append(urls[:i], urls[i+1:]...)
				c.serviceMap[service] = urls
				break
			}
		}
	}
	go func() {
		c.listenerLock.RLock()
		defer c.listenerLock.RUnlock()
		for _, l := range c.listenerMap[service] {
			l.notify(url, "down")
		}
	}()
}
