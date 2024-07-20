package config

import (
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Application App     `yaml:"app"`
		Net         Network `yaml:"network"`
		R           Redis   `yaml:"redis"`
		Log         Logger  `yaml:"log"`
		YoutubeAPI  Youtube
	}

	App struct {
		Name string `yaml:"name" env:"APP_NAME" env-required:"true"`
	}

	Network struct {
		Host string `yaml:"host" env:"APP_HOST" env-required:"true" env-default:"0.0.0.0"`
		Port string `yaml:"port" env:"APP_PORT" env-required:"true" env-default:"8888"`
	}

	Redis struct {
		Addr     string `yaml:"addr" env-required:"true"`
		Password string `env:"REDIS_PASSWORD" env-required:"true"`
	}

	Logger struct {
		Level  string `yaml:"level"`
		IsJSON bool   `yaml:"is_json"`
	}

	Youtube struct {
		YoutubeConfigFilePath string `env:"YOUTUBE_CONFIG_PATH" env-required:"true"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	err := cleanenv.ReadConfig(path.Join("./", configPath), config)
	if err != nil {
		return nil, err
	}

	err = cleanenv.UpdateEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
