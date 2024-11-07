package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DbConfig     dbConfig     `json:"db_config"`
	HttpConfig   httpConfig   `json:"http_config"`
	FolderConfig folderConfig `json:"folder_config"`
}

type dbConfig struct {
	DbHost     string `json:"db_host" env:"DB_HOST"`
	DbPort     string `json:"db_port" env:"DB_PORT"`
	DbUser     string `json:"db_user" env:"DB_USER"`
	DbPassword string `json:"db_password" env:"DB_PASSWORD"`
	DbName     string `json:"db_name" env:"DB_NAME"`
	DbSslMode  string `json:"db_sll_mode" env:"DB_SSL_MODE"`
	DbTimeZone string `json:"db_time_zone" env:"DB_TIME_ZONE"`
}
type httpConfig struct {
	HttpHost  string `json:"http_host" env:"HTTP_SERVER"`
	HttpPort  string `json:"http_port" env:"HTTP_PORT"`
	AppName   string `json:"app_name" env:"APP_NAME"`
	AppHeader string `json:"app_header" env:"APP_HEADER"`
}

type folderConfig struct {
	PublicPath string `json:"public_path" env:"PUBLIC_PATH"`
	RootPath   string `json:"root_path" env:"ROOT_PATH"`
}

func GetConfig() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("../.env", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
