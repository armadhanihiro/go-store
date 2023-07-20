package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort             string        `mapstructure:"SERVER_PORT"`
	PostgresUri            string        `mapstructure:"POSTGRES_URI"`
	LogLevel               string        `mapstructure:"LOG_LEVEL"`
	AccessTokenKey         string        `mapstructure:"ACCESS_TOKEN_KEY"`
	AccessTokenDuration    time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenKey        string        `mapstructure:"REFRESH_TOKEN_KEY"`
	RefreshTokenDuration   time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	CloudinaryCloudName    string        `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryApiKey       string        `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryApiSecret    string        `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudinaryUploadFolder string        `mapstructure:"CLOUDINARY_UPLOAD_FOLDER"`
}

func LoadConfig(fileConfigPath string) (Config, error) {
	config := Config{}

	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
