package initialize

import (
	"fmt"
	"user-service/pkg/settings"

	"github.com/spf13/viper"
)

func LoadConfig() (*settings.Config, error) {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg settings.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	return &cfg, nil
}
