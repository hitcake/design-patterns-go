package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Config struct {
	Name string
	Port int
	IP   string
}

const ConfigFileName = "config.json"

var configInstance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		data, err := os.ReadFile(ConfigFileName)
		if err != nil {
			panic(err)
		}
		configInstance = &Config{}
		err = json.Unmarshal(data, &configInstance)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Init config failed, err:", err)
				panic(err)
			} else {
				fmt.Println("Init config success")
			}
		}()
	})
	return configInstance
}

func (c *Config) String() string {
	return fmt.Sprintf("Name: %s, Port: %d, IP: %s", c.Name, c.Port, c.IP)
}
