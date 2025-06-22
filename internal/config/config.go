package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	MainConfig  MainConfig  `toml:"mainConfig"`
	MysqlConfig MysqlConfig `toml:"mysqlConfig"`
	EmailConfig EmailConfig `toml:"emailConfig"`
	LogConfig   LogConfig   `toml:"logConfig"`
}

var config *Config

func LoadConfig() error {
	if _, err := toml.DecodeFile("D:\\My Projects\\discuss\\configs\\configs.toml", &config); err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func GetConfig() *Config {
	if config == nil {
		config = new(Config)
		if err := LoadConfig(); err != nil {
			log.Fatal("Failed to load configuration:", err)
		}
	}
	return config
}

type MainConfig struct {
	AppName string `toml:"app_name"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
}

type MysqlConfig struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	DatabaseName string `toml:"databasename"`
}

type EmailConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type LogConfig struct {
	LogPath string `toml:"logPath"`
}
