package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env"`
	HTTPServer HTTPServer `yaml:"http_server"`
	Database   Database   `yaml:"database"`
	LLM        LLM        `yaml:"llm"`
	Minio      Minio      `yaml:"minio"`
	Audio      Audio      `yaml:"audio"`
	Video      Video      `yaml:"video"`
}

type HTTPServer struct {
	Address string `yaml:"address"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type LLM struct {
	Address      string `yaml:"address"`
	HistoryLimit int    `yaml:"history_limit"`
	MaxTokens    uint32 `yaml:"max_tokens"`
}

type Audio struct {
	Address string `yaml:"address"`
}

type Video struct {
	Address string `yaml:"address"`
}

type Minio struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func Load() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", configPath)
	}

	cfg := new(Config)
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, fmt.Errorf("cannot read config: %s", err)
	}

	return cfg, nil
}
