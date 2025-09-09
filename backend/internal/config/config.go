package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	Primary Primary	`koanf:"primary" validate:"required"`
	Server ServerConfig `koanf:"server" validate:"required"`
	Database DatabaseConfig `koanf:"database" validate:"required"`
	Auth AuthConfig `koanf:"auth" validate:"required"`
}

type Primary struct {
	Env string `koanf:"env" validate:"required"`
}

type ServerConfig struct {
	Port string `koanf:"port" validate:"required"`
	ReadTimeout int `koanf:"read_timeout" validate:"required"`
	WritetimeOut int `koanf:"write_timeout" validate:"required"`
	IdleTimeout int `koanf:"idle_timeout" validate:"required"`
	CORSAllowedOrigins []string `koanf:"cors_allowed_origin" validate:"required"`
}

type DatabaseConfig struct {
	Host string `koanf:"host" validate:"required"`
	Port int `koanf:"port" validate:"required"`
	User string `koanf:"user" validate:"required"`
	Password string `koanf:"password"`
	Name string `koanf:"name" validate:"required"`
	SSLMode string `koanf:"ssl_mode" validate:"required"`
	MaxOpenConns int `koanf:"max_open_conns" validate:"required"`
	MaxIdleConns int `koanf:"max_idle_conns" validate:"required"`
	ConnMaxLifeTime int `koanf:"conn_max_life_time" validate:"required"`
	ConnMaxIdleTime int `koanf:"conn_max_idle_time" validate:"required"`
}


type AuthConfig struct {
	SecretKey string `koanf:"secret_key" validate:"required"`
}

func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	
	k:= koanf.New(".")
		
}