package config

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetConfig(t *testing.T) {
	var wg sync.WaitGroup
	// 初始化10个元素的slice
	configInstanceList := make([]*Config, 10)
	//启动10个goroutine去获取配置并填充slice
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			configInstanceList[index] = GetConfig()
		}(i)
	}
	wg.Wait()
	// 比较前后相等
	for i := 1; i < 10; i++ {
		if configInstanceList[i] != configInstanceList[i-1] {
			t.Errorf("Configs are not the same")
			break
		}
	}
	// 打印其中一个
	fmt.Println(configInstanceList[0])
}
