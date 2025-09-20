package saasupgrade

import (
	"fmt"
	"sync"
)

// ContainerHistory Memento
type ContainerHistory struct {
	history map[string][]float32
	lock    sync.RWMutex
}

func NewContainerHistory() *ContainerHistory {
	return &ContainerHistory{history: make(map[string][]float32)}
}

func (c *ContainerHistory) SaveApp(tenantId string, a *App) {
	c.lock.Lock()
	defer c.lock.Unlock()
	key := fmt.Sprintf("%s_%s", tenantId, a.appId)
	if _, ok := c.history[key]; !ok {
		c.history[tenantId] = make([]float32, 0)
	}
	c.history[key] = append(c.history[key], a.appVersion)
}
func (c *ContainerHistory) GetApp(tenantId string, appId string) *App {
	c.lock.RLock()
	defer c.lock.RUnlock()
	key := fmt.Sprintf("%s_%s", tenantId, appId)
	if versionList, ok := c.history[key]; ok {

		appVersion := versionList[len(versionList)-1]
		return &App{appId: appId, appVersion: appVersion}
	}
	return nil
}
