package config

import (
	"github.com/micro/go-micro/config"
)

// Conf 配置
var Conf Config

// Permission 权限
type Permission struct {
	Service     string `json:"service"`
	Method      string `json:"method"`
	Auth        bool   `json:"auth"`
	Policy      bool   `json:"policy"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

// Config 配置结构
type Config struct {
	App         string
	Version     string
	Permissions []Permission `json:"permissions"`
}

// LoadFile 加载文件
func (c *Config) LoadFile(file string) *Config {
	config.LoadFile(file)
	config.Scan(&Conf)
	return c
}

// GetApp 获取配置
func (c *Config) GetApp() string {
	return Conf.App
}

// GetVersion 获取配置
func (c *Config) GetVersion() string {
	return Conf.Version
}

// GetPermissions 获取配置
func (c *Config) GetPermissions() []Permission {
	return Conf.Permissions
}
