package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	// Application
	GoEnv          string `mapstructure:"GO_ENV"`
	AppHost        string `mapstructure:"APP_HOST"`
	AppPort        int    `mapstructure:"APP_PORT"`
	SwaggerEnabled bool   `mapstructure:"SWAGGER_ENABLED"`

	// Database
	DbName         string `mapstructure:"DATABASE_NAME"`
	DbPort         int    `mapstructure:"DATABASE_PORT"`
	DbHost         string `mapstructure:"DATABASE_HOST"`
	DbDriver       string `mapstructure:"DATABASE_DRIVER"`
	DbUser         string `mapstructure:"DATABASE_USER"`
	DbPassword     string `mapstructure:"DATABASE_PASSWORD"`
	DbMaxOpenConns int    `mapstructure:"DATABASE_MAX_OPEN_CONNS"`
	DbMaxIdleConns int    `mapstructure:"DATABASE_MAX_IDLE_CONNS"`
	DbConnMaxLife  int    `mapstructure:"DATABASE_CONN_MAX_LIFE"`

	// Authentication
	JwtAccessTokenExpirationTime  int    `mapstructure:"JWT_ACCESS_TOKEN_EXPIRATION_TIME"`
	JwtAccessTokenSecret          string `mapstructure:"JWT_ACCESS_TOKEN_SECRET"`
	JwtRefreshTokenExpirationTime int    `mapstructure:"JWT_REFRESH_TOKEN_EXPIRATION_TIME"`
	JwtRefreshTokenSecret         string `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`

	// Redis broker
	RedisBrokerUrl string `mapstructure:"REDIS_BROKER_URL"`
}

func loadConfiguration(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Load config failed")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Err(err).Msg("Load config failed")
	}

	return config
}

var AppConfiguration = loadConfiguration(".")
