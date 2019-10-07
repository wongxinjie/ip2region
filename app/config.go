/**
* @File: config.go
* @Author: wongxinjie
* @Date: 2019/10/6
*/
package app

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	SecretKey []byte
}

func InitConfig() (*Config, error) {
	config := &Config{
		SecretKey: []byte(viper.GetString("secret_key")),
	}
	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("secret_key must be set")
	}
	return config, nil
}
