package saasupgrade

import "fmt"

// Container Originator
type Container struct {
	containerId string
	tenantId    string
	appId       string
	appVersion  float32
}

func (c *Container) SetAppVersion(appVersion float32) {
	c.appVersion = appVersion
}
func (c *Container) GetAppVersion() float32 {
	return c.appVersion
}

func (c *Container) SaveAppVersion() *App {
	return &App{appId: c.appId, appVersion: c.appVersion}
}
func (c *Container) RestoreAppVersion(a *App) {
	c.appVersion = a.appVersion
}
func (c *Container) String() string {
	return fmt.Sprintf("appId %s appVersion %0.2f", c.appId, c.appVersion)
}
