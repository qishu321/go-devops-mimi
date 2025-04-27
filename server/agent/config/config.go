package config

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

type Config struct {
	ServerURL string
	Token     string
	AgentName string
	IP        string
	Label     string
	Version   string
}

var Cfg Config

// LoadConfig 读取 config.ini 配置文件
func LoadConfig(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	Cfg = Config{
		ServerURL: cfg.Section("agent").Key("server_url").String(),
		Token:     cfg.Section("agent").Key("token").String(),
		AgentName: cfg.Section("agent").Key("agent_name").String(),
		IP:        cfg.Section("agent").Key("ip").String(),
		Label:     cfg.Section("agent").Key("label").String(),
		Version:   cfg.Section("agent").Key("version").String(),
	}

	log.Printf("加载配置: %+v\n", Cfg)
	return nil
}
