package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string        `yaml:"env" env-default:"local" env-required:"true"`
	TokenTTL   time.Duration `yaml:"token_ttl" env-default:"1h" env-required:"true"`
	HTTPServer Server        `yaml:"http_server" env-required:"true"`
	Storage    DB            `yaml:"storage" env-required:"true"`
}

type Server struct {
	Host string `yaml:"host" env-default:"localhost" env-required:"true"`
	Port int    `yaml:"port" env-default:"5432" env-required:"true"`
}

type DB struct {
	Host    string `yaml:"host" env-default:"localhost" env-required:"true"`
	Port    int    `yaml:"port" env-default:"5432" env-required:"true"`
	DBName  string `yaml:"db_name" env-required:"true"`
	User    string `yaml:"user" env-required:"true"`
	Pass    string `yaml:"pass" env-required:"true"`
	SSLMode string `yaml:"ssl_mode" env-default:"disable" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
