package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	HTTPServer HTTPServer
	Tarantool  Tarantool
	Mattermost Mattermost
}

type HTTPServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Tarantool struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true" env:"TARANTOOL_PASSWORD"`
	Space    string `yaml:"space" env-required:"true"`
}

type Mattermost struct {
	Url       string `yaml:"server_url" env-required:"true"`
	botToken  string `yaml:"bot_token" env-required:"true"`
	teamId    string `yaml:"team_id" env-required:"true"`
	channelId string `yaml:"channel_id" env-required:"true"`
}

func MustLoad() *Config {
	os.Setenv("CONFIG_PATH", "./config/config.yaml")

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		panic("config path is required, but it is empty")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic("config file does not exist" + cfgPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
