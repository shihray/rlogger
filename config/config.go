package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	setOnce      sync.Once
	globalConfig *config
)

type config struct {
	systemConfig Config
}

func NewFromViper() (*config, error) {
	var c config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func SetConfig(c *config) {
	setOnce.Do(func() {
		globalConfig = c
	})
}

func GetConfig() *config {
	return globalConfig
}
