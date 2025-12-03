package initialize

import (
	"fmt"
	"user-service/pkg/globals"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	if err := viper.Unmarshal(&globals.Config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}
}
