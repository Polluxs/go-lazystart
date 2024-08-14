package shared

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"main/logger"
)

func InitialiseAppSettings(path string) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(path)

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Error reading configuration", zap.Error(err))
	}
}
