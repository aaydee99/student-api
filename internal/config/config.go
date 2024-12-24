package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

)

type HTTPServer struct {
	addr string
}

// env_default:"production"
type Config struct {
	Env         string     `yaml:"env" env:"ENV" env_required:"true" `
	StoragePath string     `yaml:"storage_path" env_required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "config file path")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("No config file provided")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	return &cfg
}
