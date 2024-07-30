package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ECRHost   string
	AWSRegion string
	Port      string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()

	viper.SetDefault("ECR_HOST", "123456789.dkr.ecr.ap-northeast-2.amazonaws.com")
	viper.SetDefault("AWS_REGION", "ap-northeast-2")
	viper.SetDefault("PORT", "8080")

	config := &Config{
		ECRHost:   viper.GetString("ECR_HOST"),
		AWSRegion: viper.GetString("AWS_REGION"),
		Port:      viper.GetString("PORT"),
	}

	log.Printf("Config loaded: %+v\n", config)
	return config
}
