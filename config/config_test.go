package config_test

import (
	"fmt"
	"shorten/config"
	"testing"
)

func TestConfigInstance(t *testing.T) {
	conf := config.GetInstance("../etc/conf/config.json")
	if conf == nil {
		t.Error("Config instance null")
	}
	fmt.Println(t.Name() + ":OK")
}
