package config

import (
	"encoding/json"
	"os"
	"sync"
)

var instance *Config
var once sync.Once

type Config struct {
	DBtype     string `json:"db_type"`
	ServerInfo `json:"server_info"`
}

type ServerInfo struct {
	Port     int        `json:"port"`
	APIroute []APIroute `json:"api_route"`
	Static   []Static   `json:"static"`
}

type APIroute struct {
	Path    string `json:"path"`
	Method  string `json:"method"`
	Handler string `json:"handler"`
}

type Static struct {
	Path string `json:"path"`
	File string `json:"file"`
}

func GetInstance(path ...string) *Config {
	once.Do(func() {
		if len(path) > 0 {
			f, err := os.Open(path[0])
			if err != nil {
				panic("Loading config error")
			}
			instance = &Config{}
			err = json.NewDecoder(f).Decode(instance)
			if err != nil {
				panic(err)
			}
		}
	})
	return instance
}
