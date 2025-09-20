package saasupgrade

import (
	"fmt"
	"testing"
)

func TestUpgrade(t *testing.T) {
	app := &App{appId: "Customer", appVersion: 1.0}
	container := &Container{containerId: "123123", tenantId: "1000001", appId: app.appId, appVersion: app.appVersion}
	fmt.Println("初始化版本" + container.String())
	// upgrade
	containerHistory := NewContainerHistory()
	containerHistory.SaveApp(container.tenantId, app)
	app2 := &App{appId: "Customer", appVersion: 2.0}
	container.SetAppVersion(app2.appVersion)
	fmt.Println("升级后版本" + container.String())
	// 出现异常，需要回滚
	a := containerHistory.GetApp(container.tenantId, app2.appId)
	if a != nil {
		container.RestoreAppVersion(a)
	}
	fmt.Println("回滚后后版本" + container.String())
}
