/**
* @File: config.go
* @Author: wongxinjie
* @Date: 2019/10/6
 */
package api

import "github.com/spf13/viper"

type Config struct {
	Host       string
	Port       int
	ProxyCount int
}

func InitConfig() (*Config, error) {
	config := &Config{
		Host:       viper.GetString("host"),
		Port:       viper.GetInt("port"),
		ProxyCount: viper.GetInt("proxy_count"),
	}
	if config.Port == 0 {
		config.Port = 12000
	}
	return config, nil
}
