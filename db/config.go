/**
* @File: config.go
* @Author: wongxinjie
* @Date: 2019/10/6
 */
package db

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURI string
}

func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("database_uri"),
	}
	if config.DatabaseURI == "" {
		return nil, fmt.Errorf("database_uri must be set")
	}
	return config, nil
}
